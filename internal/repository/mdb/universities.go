package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models/university"
)

// UniversityRepo - Структура для работы с коллекцией из mongoDB
type UniversityRepo struct {
	db *mongo.Collection
}

// NewUniversityRepo - Создание репозитория для работы с коллекцией
func NewUniversityRepo(db *mongo.Database) *UniversityRepo {
	return &UniversityRepo{
		db: db.Collection(universitiesCollection),
	}
}

// GetByDomain - Получение записи об университете по имени домена
func (r *UniversityRepo) GetByDomain(ctx context.Context, domainName string) (university.University, error) {
	var univ university.University
	err := r.db.FindOne(ctx, bson.M{
		"domain": domainName,
	}).Decode(&univ)

	return univ, err
}
