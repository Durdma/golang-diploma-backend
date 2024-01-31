package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models/university"
)

type UniversityRepo struct {
	db *mongo.Collection
}

func NewUniversityRepo(db *mongo.Database) *UniversityRepo {
	return &UniversityRepo{
		db: db.Collection(universitiesCollection),
	}
}

func (r *UniversityRepo) GetByDomain(ctx context.Context, domainName string) (university.University, error) {
	var univ university.University
	err := r.db.FindOne(ctx, bson.M{
		"domain": domainName,
	}).Decode(&univ)

	return univ, err
}
