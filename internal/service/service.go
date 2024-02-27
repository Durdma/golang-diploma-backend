package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"sas/pkg/auth"
	"sas/pkg/cache"
	"sas/pkg/email"
	"sas/pkg/hash"
	"time"
)

// Universities - Интерфейс для взаимодействия с сервисом университетов
type Universities interface {
	GetByDomain(ctx context.Context, domainName string) (models.University, error)
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type AdminSignUpInput struct {
	Name     string
	Email    string
	Password string
	Domain   string
}

type AdminSignInInput struct {
	Email    string
	Password string
	Domain   string
}

type Admins interface {
	SignIn(ctx context.Context, input AdminSignInInput) (Tokens, error)
	SignUp(ctx context.Context, input AdminSignUpInput) error
	RefreshTokens(ctx context.Context, domain primitive.ObjectID, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, hash string) error
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
	SignIn(ctx context.Context, input EditorSignInInput) (Tokens, error)
	SignUp(ctx context.Context, input EditorSignUpInput) error
	RefreshTokens(ctx context.Context, university primitive.ObjectID, refreshToken string) (Tokens, error)
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
func NewServices(repos *repository.Repositories, cache cache.Cache, hasher hash.PasswordHasher,
	tokenManager auth.TokenManager, emailProvider email.Provider, accessTTL time.Duration,
	refreshTTL time.Duration) *Services {
	emailService := NewEmailsService(emailProvider, "")
	return &Services{
		Universities: NewUniversitiesService(repos.Universities, cache),
		Editors:      NewEditorsService(repos.Editors, hasher, tokenManager, emailService, accessTTL, refreshTTL),
	}
}
