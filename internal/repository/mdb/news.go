package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
)

type NewsRepo struct {
	db *mongo.Collection
}

func NewNewsRepo(db *mongo.Database) *NewsRepo {
	return &NewsRepo{
		db: db.Collection(newsCollection),
	}
}

func (r *NewsRepo) Create(ctx context.Context, news models.News) error {
	_, err := r.db.InsertOne(ctx, news)
	return err
}

func (r *NewsRepo) Update(ctx context.Context, news models.News) error {
	return nil
}

func (r *NewsRepo) Delete(ctx context.Context, newsId primitive.ObjectID) error {
	return nil
}
