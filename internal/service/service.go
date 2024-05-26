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
	AddUniversity(ctx context.Context, domainId primitive.ObjectID, domainName string, shortName string) (primitive.ObjectID, error)
	GetByDomain(ctx context.Context, domainName string) (models.University, error)
	GetByUniversityId(ctx context.Context, universityIdStr primitive.ObjectID) (models.University, error)
	GetUniversityColors(ctx context.Context, universityId primitive.ObjectID) (map[string]string, error)
	PatchUniversityCSS(ctx context.Context, universityId primitive.ObjectID, colors map[string]string) error
	SetUniversityHistory(ctx context.Context, universityId primitive.ObjectID, history models.History) error
}

type NewsInput struct {
	Header       string
	Description  string
	Body         string
	CreatedBy    string
	UniversityId primitive.ObjectID
}

type News interface {
	AddNews(ctx context.Context, input NewsInput) (primitive.ObjectID, error)
	GetAllNews(ctx context.Context, domainId primitive.ObjectID) ([]models.News, error)
	AddHeaderImageURL(ctx context.Context, recordId string, imageURL string) error
}

type DocsInput struct {
	UniversityId    primitive.ObjectID
	Header          string
	Description     string
	Code            string
	Magistrate      bool
	Enrollee        bool
	PublicationDate time.Time
	CreatedBy       string
}

type Docs interface {
	AddDocs(ctx context.Context, input DocsInput) (primitive.ObjectID, error)
	GetAllUniversityDocs(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error)
	GetAllBachelors(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error)
	GetAllMags(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error)
	GetAllEnrollsDocs(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error)
	AddDocsURL(ctx context.Context, docId string, docURL string) error
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

type Users interface {
	SignUp(ctx context.Context) error
	SignIn(ctx context.Context, input SignInInput) (models.User, Tokens, error)
	RefreshTokens(ctx context.Context, domain primitive.ObjectID, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, domain primitive.ObjectID, hash string) error
	GetUserById(ctx context.Context, userId string) (models.User, error)
}

type Admins interface {
	SignUp(ctx context.Context, input AdminSignUpInput) error
	GetAll(ctx context.Context) ([]models.User, error)
	ChangeAdminBlockStatus(ctx context.Context, editorId string, state string) error
	ChangeAdminVerifyStatus(ctx context.Context, editorId string, state string) error
	UpdateAdmin(ctx context.Context, newUser UpdateAdminInput) error
}

type DomainInput struct {
	DomainName string
	ShortName  string
	HTTPName   string
	Verify     bool
	Visible    bool
}

type Domains interface {
	AddDomain(ctx context.Context, input DomainInput) (primitive.ObjectID, error)
	DeleteDomain(ctx context.Context, domainId primitive.ObjectID) error
	GetByHTTPName(ctx context.Context, domain string) (models.Domain, error)
	GetById(ctx context.Context, domainId primitive.ObjectID) (models.Domain, error)
	GetByDomainName(ctx context.Context, domainName string) (models.Domain, error)
	GetAllDomains(ctx context.Context) ([]models.Domain, error)
	ChangeSiteVisibleStatus(ctx context.Context, domainId string, state string) error
	ChangeSiteVerifyStatus(ctx context.Context, domainId string, state string) error
	UpdateDomain(ctx context.Context, newDomain UpdateDomainInput) error
	AddUniversityId(ctx context.Context, domainId primitive.ObjectID, universityId primitive.ObjectID) error
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
	News         News
	Docs         Docs
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
		News:         NewNewsService(repos.News),
		Docs:         NewDocsService(repos.Docs),
	}
}
