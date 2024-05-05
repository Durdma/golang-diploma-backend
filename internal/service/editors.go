package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"sas/pkg/auth"
	"sas/pkg/hash"
	"sas/pkg/logger"
	"time"
)

// EditorsService - Структура сервиса редакторов
type EditorsService struct {
	repo         repository.Editors
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	emailService Emails

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

// NewEditorsService - Создание нового сервиса редакторов
func NewEditorsService(repo repository.Editors, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	emailService Emails, accessTTL time.Duration, refreshTTL time.Duration) *EditorsService {
	return &EditorsService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		emailService:    emailService,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

// SignUp - Регистрация нового редактора на сайте университета
func (s *EditorsService) SignUp(ctx context.Context, input EditorSignUpInput) error {
	verificationCode := primitive.NewObjectID()
	editor := models.Editor{
		Name:         input.Name,
		Password:     s.hasher.Hash(input.Password),
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
		UniversityID: input.UniversityID,
		Verification: models.Verification{
			Code: verificationCode,
		},
	}

	if err := s.repo.Create(ctx, editor); err != nil {
		return err
	}

	return s.emailService.AddToList(AddToListInput{
		Email:            input.Email,
		Name:             input.Name,
		VerificationCode: verificationCode.Hex(),
	})
}

// SignIn - Вход редактора на сайт по паре логин-пароль
func (s *EditorsService) SignIn(ctx context.Context, input EditorSignInInput) (Tokens, error) {
	editor, err := s.repo.GetByCredentials(ctx, input.UniversityID, input.Email, s.hasher.Hash(input.Password))
	logger.Info(editor)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(ctx, editor.ID)
}

func (s *EditorsService) RefreshTokens(ctx context.Context, universityId primitive.ObjectID, refreshToken string) (Tokens, error) {
	editor, err := s.repo.GetByRefreshToken(ctx, universityId, refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(ctx, editor.ID)
}

// Verify - Подтверждение регистрации нового редактора
func (s *EditorsService) Verify(ctx context.Context, code string) error {
	return s.repo.Verify(ctx, code)
}

func (s *EditorsService) createSession(ctx context.Context, editorId primitive.ObjectID) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessTokenTTL, res.AccessToken, err = s.tokenManager.NewJWT(editorId.Hex(), s.accessTokenTTL)
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
	err = s.repo.SetSession(ctx, editorId, session)

	logger.Info("session added")

	return res, err
}
