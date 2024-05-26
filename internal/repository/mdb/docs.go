package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sas/internal/models"
)

type DocsRepo struct {
	db *mongo.Collection
}

func NewDocsRepo(db *mongo.Database) *DocsRepo {
	return &DocsRepo{db: db.Collection(docsCollection)}
}

func (r *DocsRepo) Create(ctx context.Context, docs models.Docs) (primitive.ObjectID, error) {
	docId, err := r.db.InsertOne(ctx, docs)

	return docId.InsertedID.(primitive.ObjectID), err
}

func (r *DocsRepo) AddDocsURL(ctx context.Context, docId primitive.ObjectID, docURL string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": docId,
	},
		bson.M{
			"$set": bson.M{"doc_url": docURL},
		})

	return err
}

func (r *DocsRepo) GetAllUniversityDocs(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error) {
	var docs []models.Docs

	filter := bson.M{
		"university_id": universityId,
	}

	sort := bson.D{{"created_at", 1}}
	opts := options.Find().SetSort(sort)
	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs, err
}

func (r *DocsRepo) GetAllBachelors(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error) {
	var docs []models.Docs

	filter := bson.M{
		"university_id": universityId,
		"code":          bson.M{"$regex": ".*", "$options": "i"},
		"magistrate":    false,
	}

	sort := bson.D{{"code", 1}}
	opts := options.Find().SetSort(sort)
	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs, err
}

func (r *DocsRepo) GetAllMags(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error) {
	var docs []models.Docs

	filter := bson.M{
		"university_id": universityId,
		"code":          bson.M{"$regex": ".*", "$options": "i"},
		"magistrate":    true,
	}

	sort := bson.D{{"code", 1}}
	opts := options.Find().SetSort(sort)
	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs, err
}

func (r *DocsRepo) GetAllEnrollsDocs(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error) {
	var docs []models.Docs

	filter := bson.M{
		"university_id": universityId,
		"code":          "",
		"enrollee":      true,
	}

	sort := bson.D{{"created_at", 1}}
	opts := options.Find().SetSort(sort)
	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs, err
}
