package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models/university"
	"sas/internal/repository"
	"sas/pkg/hash"
	"time"
)

// EditorsService - Структура сервиса редакторов
type EditorsService struct {
	repo   repository.Editors
	hasher hash.PasswordHasher
}

// NewEditorsService - Создание нового сервиса редакторов
func NewEditorsService(repo repository.Editors, hasher hash.PasswordHasher) *EditorsService {
	return &EditorsService{
		repo:   repo,
		hasher: hasher,
	}
}

// SignIn - Вход редактора на сайт по паре логин-пароль
func (s *EditorsService) SignIn(ctx context.Context, email string, password string) (string, error) {
	return "", nil
}

// SignUp - Регистрация нового редактора на сайте университета
func (s *EditorsService) SignUp(ctx context.Context, input EditorSignUpInput) error {
	editor := university.Editor{
		Name:         input.Name,
		Password:     s.hasher.Hash(input.Password),
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
		UniversityID: input.UniversityID,
		Verification: university.Verification{
			Hash: primitive.NewObjectID(),
		},
	}

	return s.repo.Create(ctx, editor)
}

// Verify - Подтверждение регистрации нового редактора
func (s *EditorsService) Verify(ctx context.Context, hash string) error {
	return s.repo.Verify(ctx, hash)
}
