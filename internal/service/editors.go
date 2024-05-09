package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"sas/pkg/auth"
	"sas/pkg/hash"
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

func (s *EditorsService) ChangeEditorBlockStatus(ctx context.Context, editorId string, state string) error {
	stateBool := false
	if state == "true" {
		stateBool = true
	} else {
		if state == "false" {
			stateBool = false
		} else {
			return errors.New("incorrect state")
		}
	}

	return s.repo.ChangeBlockStatus(ctx, editorId, stateBool)
}

func (s *EditorsService) ChangeEditorVerifyStatus(ctx context.Context, editorId string, state string) error {
	stateBool := false
	if state == "true" {
		stateBool = true
	} else {
		if state == "false" {
			stateBool = false
		} else {
			return errors.New("incorrect state")
		}
	}

	return s.repo.ChangeVerificationStatus(ctx, editorId, stateBool)
}

func (s *EditorsService) ChangeEditor(ctx context.Context) error {
	return nil
}
