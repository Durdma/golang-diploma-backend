package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models/university"
	"sas/internal/repository/mdb"
)

// Universities - Интерфейс для репозитория университетов
type Universities interface {
	GetByDomain(ctx context.Context, domain string) (university.University, error)
}

// Editors - Интерфейс для репозитория редакторов
type Editors interface {
	Create(ctx context.Context, editor university.Editor) error
	GetByCredentials(ctx context.Context, email, password university.Editor) error
	Verify(ctx context.Context, code string) error
}

// Repositories - структура со всеми репозиториями
type Repositories struct {
	Universities Universities
	Editors      Editors
}

// NewRepositories - Создание общего репозитория
func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Universities: mdb.NewUniversityRepo(db),
		Editors:      mdb.NewEditorsRepo(db),
	}
}
