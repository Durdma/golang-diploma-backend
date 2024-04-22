package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"sas/pkg/auth"
	"sas/pkg/hash"
	"time"
)

type AdminsService struct {
	repo         repository.Admins
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	emailService Emails

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAdminsService(repo repository.Admins, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	emailService Emails, accessTTL time.Duration, refreshTTL time.Duration) *AdminsService {
	return &AdminsService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		emailService:    emailService,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

func (s *AdminsService) SignUp(ctx context.Context, input AdminSignUpInput) error {
	verificationCode := primitive.NewObjectID()
	adm := models.Admin{
		Name:         input.Name,
		Email:        input.Email,
		Password:     s.hasher.Hash(input.Password),
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
		Verification: models.Verification{
			Code: verificationCode,
		},
	}

	if err := s.repo.Create(ctx, adm); err != nil {
		return err
	}

	return s.emailService.AddToListAdmin(AddToListInput{
		Email:            input.Email,
		Name:             input.Name,
		VerificationCode: verificationCode.Hex(),
	})
}

func (s *AdminsService) SignIn(ctx context.Context, input AdminSignInInput) (Tokens, error) {
	admin, err := s.repo.GetByCredentials(ctx, input.Email, s.hasher.Hash(input.Password))
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(ctx, admin.ID)
}

func (s *AdminsService) RefreshTokens(ctx context.Context, domain primitive.ObjectID, refreshToken string) (Tokens, error) {
	admin, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(ctx, admin.ID)
}

func (s *AdminsService) Verify(ctx context.Context, hash string) error {
	return s.repo.Verify(ctx, hash)
}

func (s *AdminsService) createSession(ctx context.Context, adminId primitive.ObjectID) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(adminId.Hex(), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := models.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err = s.repo.SetSession(ctx, adminId, session)

	return res, err
}
