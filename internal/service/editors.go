package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"sas/pkg/hash"
	"time"
)

// EditorsService - Структура сервиса редакторов
type EditorsService struct {
	repo         repository.Editors
	hasher       hash.PasswordHasher
	emailService Emails
}

// NewEditorsService - Создание нового сервиса редакторов
func NewEditorsService(repo repository.Editors, hasher hash.PasswordHasher, emailService Emails) *EditorsService {
	return &EditorsService{
		repo:         repo,
		hasher:       hasher,
		emailService: emailService,
	}
}

// SignIn - Вход редактора на сайт по паре логин-пароль
func (s *EditorsService) SignIn(ctx context.Context, email string, password string) (string, error) {
	return "", nil
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

// Verify - Подтверждение регистрации нового редактора
func (s *EditorsService) Verify(ctx context.Context, code string) error {
	return s.repo.Verify(ctx, code)
}
