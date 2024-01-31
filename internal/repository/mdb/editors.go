// Package mdb Editors реализуются пока как студенты для понимания сути программы
// потом переписать под admins из примера
// TODO потом переписать логику на админов из примера
package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models/university"
)

type EditorsRepo struct {
	db *mongo.Collection
}

func NewEditorsRepo(db *mongo.Database) *EditorsRepo {
	return &EditorsRepo{
		db: db.Collection(editorsCollection),
	}
}

func (r *EditorsRepo) Create(ctx context.Context, editor university.Editor) error {
	_, err := r.db.InsertOne(ctx, editor)
	return err
}

func (r *EditorsRepo) GetByCredentials(ctx context.Context, email, password university.Editor) error {
	return nil
}

func (r *EditorsRepo) Verify(ctx context.Context, hash string) error {
	return nil
}
