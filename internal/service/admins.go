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
