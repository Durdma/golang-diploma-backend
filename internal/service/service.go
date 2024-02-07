package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models/university"
	"sas/internal/repository"
	"sas/pkg/cache"
	"sas/pkg/email"
	"sas/pkg/hash"
)

// Universities - Интерфейс для взаимодействия с сервисом университетов
type Universities interface {
	GetByDomain(ctx context.Context, domainName string) (university.University, error)
}

// EditorSignUpInput TODO Взято из примера для понимания, при добавлении редакторов переписать
// EditorSignUpInput - Структура для парсинга данных при регистрации в go-объект из json
type EditorSignUpInput struct {
	Name         string
	Email        string
	Password     string
	UniversityID primitive.ObjectID
}

type EditorSignInInput struct {
	Email        string
	Password     string
	UniversityID primitive.ObjectID
}

// Editors - Интерфейс для сервиса редакторов
type Editors interface {
	SignIn(ctx context.Context, email string, password string) (string, error)
	SignUp(ctx context.Context, input EditorSignUpInput) error
	Verify(ctx context.Context, hash string) error
}

type AddToListInput struct {
	Email            string
	Name             string
	VerificationCode string
}

type Emails interface {
	AddToList(input AddToListInput) error
}

// Services - Объединение всех сервисов
type Services struct {
	Universities Universities
	Editors      Editors
}

// NewServices - Создание нового сервиса
func NewServices(repos *repository.Repositories, cache cache.Cache, hasher hash.PasswordHasher, emailProvider email.Provider) *Services {
	emailService := NewEmailsService(emailProvider, "")
	return &Services{
		Universities: NewUniversitiesService(repos.Universities, cache),
		Editors:      NewEditorsService(repos.Editors, hasher, emailService),
	}
}
