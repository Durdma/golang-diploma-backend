package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
)

type AdminsRepo struct {
	db *mongo.Collection
}

func NewAdminsRepo(db *mongo.Database) *AdminsRepo {
	return &AdminsRepo{
		db: db.Collection(adminsCollection),
	}
}

func (r *AdminsRepo) Create(ctx context.Context, adm models.Admin) error {
	_, err := r.db.InsertOne(ctx, adm)
	return err
}
