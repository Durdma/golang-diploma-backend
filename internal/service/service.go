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
	AccessToken     string
	RefreshToken    string
	AccessTokenTTL  int
	RefreshTokenTTL int
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

type Users interface {
	SignUp(ctx context.Context) error
	SignIn(ctx context.Context, input SignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, domain primitive.ObjectID, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, domain primitive.ObjectID, hash string) error
}

type Admins interface {
	SignIn(ctx context.Context, input AdminSignInInput) (Tokens, error)
	SignUp(ctx context.Context, input AdminSignUpInput) error
	RefreshTokens(ctx context.Context, domain primitive.ObjectID, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, hash string) error
}

type DomainInput struct {
	HTTPDomain string
}

type Domains interface {
	AddDomain(ctx context.Context, input DomainInput) error
	GetAllDomains(ctx context.Context) ([]models.Domain, error)
	GetDomain(ctx context.Context, input string) (models.Domain, error)
	DeleteDomain(ctx context.Context, input string) error
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
	AddToListAdmin(input AddToListInput) error
}

// Services - Объединение всех сервисов
type Services struct {
	Admins       Admins
	Universities Universities
	Editors      Editors
	Domains      Domains
	Users        Users
}

// NewServices - Создание нового сервиса
func NewServices(repos *repository.Repositories, cache cache.Cache, hasher hash.PasswordHasher,
	tokenManager auth.TokenManager, emailProvider email.Provider, accessTTL time.Duration,
	refreshTTL time.Duration) *Services {
	emailService := NewEmailsService(emailProvider, "")
	return &Services{
		Admins:       NewAdminsService(repos.Admins, hasher, tokenManager, emailService, accessTTL, refreshTTL),
		Universities: NewUniversitiesService(repos.Universities, cache),
		Editors:      NewEditorsService(repos.Editors, hasher, tokenManager, emailService, accessTTL, refreshTTL),
		Domains:      NewDomainsService(repos.DNS),
		Users:        NewUsersService(repos.Users, hasher, tokenManager, emailService, accessTTL, refreshTTL),
	}
}
