package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *NewsRepo) Create(ctx context.Context, news models.News) (primitive.ObjectID, error) {
	res, err := r.db.InsertOne(ctx, news)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), err
}

func (r *NewsRepo) GetAllNews(ctx context.Context, domainId primitive.ObjectID) ([]models.News, error) {
	var news []models.News

	filter := bson.M{
		"university_id": domainId,
		"published":     true,
	}

	sort := bson.D{{"created_at", 1}}

	opts := options.Find().SetSort(sort)
	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &news); err != nil {
		return nil, err
	}

	return news, err
}

func (r *NewsRepo) AddHeaderImageURL(ctx context.Context, recordId primitive.ObjectID, imageURL string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": recordId,
	},
		bson.M{
			"$set": bson.M{"image_url": imageURL},
		})

	return err
}

func (r *NewsRepo) Update(ctx context.Context, news models.News) error {
	return nil
}

func (r *NewsRepo) Delete(ctx context.Context, newsId primitive.ObjectID) error {
	return nil
}
