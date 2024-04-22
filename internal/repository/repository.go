package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
	"sas/internal/repository/mdb"
)

// Universities - Интерфейс для репозитория университетов
type Universities interface {
	GetByDomain(ctx context.Context, domain string) (models.University, error)
}

type Admins interface {
	Create(ctx context.Context, adm models.Admin) error
	GetByCredentials(ctx context.Context, email string, password string) (models.Admin, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (models.Admin, error)
	SetSession(ctx context.Context, userId primitive.ObjectID, session models.Session) error
	Verify(ctx context.Context, code string) error
}

// Editors - Интерфейс для репозитория редакторов
type Editors interface {
	Create(ctx context.Context, editor models.Editor) error
	GetByCredentials(ctx context.Context, universityId primitive.ObjectID, email string, password string) (models.Editor, error)
	GetByRefreshToken(ctx context.Context, universityId primitive.ObjectID, refreshToken string) (models.Editor, error)
	SetSession(ctx context.Context, userId primitive.ObjectID, session models.Session) error
	Verify(ctx context.Context, code string) error
}

type News interface {
	Create(ctx context.Context, news models.News) error
	Update(ctx context.Context, news models.News) error
	Delete(ctx context.Context, newsId primitive.ObjectID) error
}

// Repositories - структура со всеми репозиториями
type Repositories struct {
	Admins       Admins
	Universities Universities
	Editors      Editors
	News         News
}

// NewRepositories - Создание общего репозитория
func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Admins:       mdb.NewAdminsRepo(db),
		Universities: mdb.NewUniversityRepo(db),
		Editors:      mdb.NewEditorsRepo(db),
		News:         mdb.NewNewsRepo(db),
	}
}
