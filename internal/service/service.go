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
	SignIn(ctx context.Context, input SignInInput) (models.User, Tokens, error)
	RefreshTokens(ctx context.Context, domain primitive.ObjectID, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, domain primitive.ObjectID, hash string) error
	GetUserById(ctx context.Context, userId string) (models.User, error)
}

type SiteInput struct {
	Name           string
	ShortName      string
	HTTPDomainName string
}

type Sites interface {
	AddNewSite(ctx context.Context, input SiteInput) (primitive.ObjectID, error)
}

type Admins interface {
	SignUp(ctx context.Context, input AdminSignUpInput) error
}

type DomainInput struct {
	HTTPDomain string
	SiteId     primitive.ObjectID
	Name       string
	ShortName  string
	Verified   bool
}

type Domains interface {
	AddDomain(ctx context.Context, input DomainInput) error
	DeleteDomain(ctx context.Context, domainId primitive.ObjectID) error
	GetByHTTPName(ctx context.Context, domain string) (models.Domain, error)
	GetById(ctx context.Context, domainId primitive.ObjectID) (models.Domain, error)
	GetByDomainName(ctx context.Context, domainName string) (models.Domain, error)
	GetAllDomains(ctx context.Context) ([]models.Domain, error)
}

// EditorSignUpInput TODO Взято из примера для понимания, при добавлении редакторов переписать
// EditorSignUpInput - Структура для парсинга данных при регистрации в go-объект из json
type EditorSignUpInput struct {
	Name       string
	Email      string
	Password   string
	DomainName string
	DomainId   primitive.ObjectID
	Verify     bool
	Block      bool
}

// Editors - Интерфейс для сервиса редакторов
type Editors interface {
	SignUp(ctx context.Context, input EditorSignUpInput) error
	ChangeEditorBlockStatus(ctx context.Context, editorId string, state string) error
	ChangeEditorVerifyStatus(ctx context.Context, editorId string, state string) error
	GetAllEditors(ctx context.Context) ([]models.User, error)
	UpdateEditor(ctx context.Context, newUser UpdateEditorInput) error
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
	Sites        Sites
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
		Sites:        NewSitesService(repos.Sites),
	}
}
