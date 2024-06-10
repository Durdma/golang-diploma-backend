package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
	"sas/internal/repository/mdb"
)

type Domains interface {
	Create(ctx context.Context, domain models.Domain) (primitive.ObjectID, error)
	Delete(ctx context.Context, domain primitive.ObjectID) error
	GetByHTTPName(ctx context.Context, domain string) (models.Domain, error)
	GetById(ctx context.Context, domainId primitive.ObjectID) (models.Domain, error)
	GetByDomainName(ctx context.Context, domainName string) (models.Domain, error)
	GetAllDomains(ctx context.Context) ([]models.Domain, error)
	ChangeVisibleStatus(ctx context.Context, domainId string, state bool) error
	ChangeVerificationStatus(ctx context.Context, domainId string, state bool) error
	UpdateDomain(ctx context.Context, domain models.Domain) error
	AddUniversityId(ctx context.Context, domainId primitive.ObjectID, universityId primitive.ObjectID) error
}

// Universities - Интерфейс для репозитория университетов
type Universities interface {
	Create(ctx context.Context, university models.University) (primitive.ObjectID, error)
	GetByDomain(ctx context.Context, domain string) (models.University, error)
	GetByUniversityId(ctx context.Context, sitId primitive.ObjectID) (models.University, error)
	ChangeCSS(ctx context.Context, universityId primitive.ObjectID, colors map[string]string) error
	SetUniversityHistory(ctx context.Context, universityId primitive.ObjectID, history models.History) error
}

type Users interface {
	Create(ctx context.Context, user models.User) error
	GetByCredentials(ctx context.Context, email string, password string, domain primitive.ObjectID) (models.User, error)
	GetByRefreshToken(ctx context.Context, domain primitive.ObjectID, refreshToken string) (models.User, error)
	GetUserById(ctx context.Context, userId primitive.ObjectID) (models.User, error)
	SetSession(ctx context.Context, userId primitive.ObjectID, session models.Session) error
	Verify(ctx context.Context, domain primitive.ObjectID, code string) error
}

type Admins interface {
	Create(ctx context.Context, adm models.User) error
	GetAll(ctx context.Context) ([]models.User, error)
	ChangeBlockStatus(ctx context.Context, adminId string, state bool) error
	ChangeVerificationStatus(ctx context.Context, editorId string, state bool) error
	UpdateAdmin(ctx context.Context, user models.User) error
	GetAdminById(ctx context.Context, adminId primitive.ObjectID) (models.User, error)
}

// Editors - Интерфейс для репозитория редакторов
type Editors interface {
	Create(ctx context.Context, editor models.User) error
	ChangeBlockStatus(ctx context.Context, editorId string, state bool) error
	ChangeVerificationStatus(ctx context.Context, editorId string, state bool) error
	GetEditorById(ctx context.Context, userId primitive.ObjectID) (models.User, error)
	UpdateEditor(ctx context.Context, user models.User) error
	GetAllEditors(ctx context.Context) ([]models.User, error)
}

type News interface {
	Create(ctx context.Context, news models.News) (primitive.ObjectID, error)
	Update(ctx context.Context, news models.News) error
	Delete(ctx context.Context, newsId primitive.ObjectID) error
	GetAllNews(ctx context.Context, domainId primitive.ObjectID) ([]models.News, error)
	AddHeaderImageURL(ctx context.Context, recordId primitive.ObjectID, imageURL string) error
}

type Docs interface {
	Create(ctx context.Context, docs models.Docs) (primitive.ObjectID, error)
	GetAllUniversityDocs(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error)
	GetAllBachelors(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error)
	GetAllMags(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error)
	GetAllEnrollsDocs(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error)
	AddDocsURL(ctx context.Context, docId primitive.ObjectID, docURL string) error
}

// Repositories - структура со всеми репозиториями
type Repositories struct {
	Admins       Admins
	Universities Universities
	Editors      Editors
	News         News
	DNS          Domains
	Users        Users
	Docs         Docs
}

// NewRepositories - Создание общего репозитория
func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Admins:       mdb.NewAdminsRepo(db),
		Universities: mdb.NewUniversityRepo(db),
		Editors:      mdb.NewEditorsRepo(db),
		News:         mdb.NewNewsRepo(db),
		DNS:          mdb.NewDNSRepo(db),
		Users:        mdb.NewUsersRepo(db),
		Docs:         mdb.NewDocsRepo(db),
	}
}
