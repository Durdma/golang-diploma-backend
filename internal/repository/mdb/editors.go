// Package mdb Editors реализуются пока как студенты для понимания сути программы
// потом переписать под admins из примера
// TODO потом переписать логику на админов из примера
package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models/university"
)

// EditorsRepo - Структура для работы с коллекцией из mongoDB
type EditorsRepo struct {
	db *mongo.Collection
}

// NewEditorsRepo - Создание нового репозитория
func NewEditorsRepo(db *mongo.Database) *EditorsRepo {
	return &EditorsRepo{
		db: db.Collection(editorsCollection),
	}
}

// Create - Добавление записи о новом редакторе в коллекцию
func (r *EditorsRepo) Create(ctx context.Context, editor university.Editor) error {
	_, err := r.db.InsertOne(ctx, editor)
	return err
}

// GetByCredentials - Получение записи из коллекции по определенным полям
func (r *EditorsRepo) GetByCredentials(ctx context.Context, email, password university.Editor) error {
	return nil
}

// Verify - Подтверждение учетной зарегистрированной учетной записи
func (r *EditorsRepo) Verify(ctx context.Context, hash string) error {
	hashId, err := primitive.ObjectIDFromHex(hash)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx,
		bson.M{"verification.hash": hashId},
		bson.M{"$set": bson.M{"verification.verified": true}})

	return err
}
