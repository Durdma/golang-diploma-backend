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

type UsersService struct {
	repo         repository.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	emailService Emails

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	emailService Emails, accessTTL time.Duration, refreshTTL time.Duration) *UsersService {
	return &UsersService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		emailService:    emailService,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

func (s *UsersService) SignUp(ctx context.Context) error {
	return nil
}

type SignInInput struct {
	Email    string
	Password string
	Domain   primitive.ObjectID
}

func (s *UsersService) SignIn(ctx context.Context, input SignInInput) (models.User, Tokens, error) {
	user, err := s.repo.GetByCredentials(ctx, input.Email, s.hasher.Hash(input.Password), input.Domain)
	if err != nil {
		return models.User{}, Tokens{}, err
	}

	session, err := s.createSession(ctx, user.ID)
	if err != nil {
		return models.User{}, Tokens{}, err
	}

	return user, session, err
}

func (s *UsersService) GetUserById(ctx context.Context, userId string) (models.User, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return models.User{}, err
	}

	return s.repo.GetUserById(ctx, id)
}

func (s *UsersService) RefreshTokens(ctx context.Context, domain primitive.ObjectID, refreshToken string) (Tokens, error) {
	user, err := s.repo.GetByRefreshToken(ctx, domain, refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(ctx, user.ID)
}

func (s *UsersService) Verify(ctx context.Context, domain primitive.ObjectID, hash string) error {
	return s.repo.Verify(ctx, domain, hash)
}

func (s *UsersService) createSession(ctx context.Context, userId primitive.ObjectID) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessTokenTTL, res.AccessToken, err = s.tokenManager.NewJWT(userId.Hex(), s.accessTokenTTL)
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

	err = s.repo.SetSession(ctx, userId, session)

	res.RefreshTokenTTL = int(s.refreshTokenTTL.Seconds())

	return res, err
}
