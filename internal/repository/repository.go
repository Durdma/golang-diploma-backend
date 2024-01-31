package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models/university"
	"sas/internal/repository/mdb"
)

type Universities interface {
	GetByDomain(ctx context.Context, domain string) (university.University, error)
}

type Editors interface {
	Create(ctx context.Context, editor university.Editor) error
	GetByCredentials(ctx context.Context, email, password university.Editor) error
	Verify(ctx context.Context, hash string) error
}

type Repositories struct {
	Universities Universities
	Editors      Editors
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Universities: mdb.NewUniversityRepo(db),
		Editors:      mdb.NewEditorsRepo(db),
	}
}
