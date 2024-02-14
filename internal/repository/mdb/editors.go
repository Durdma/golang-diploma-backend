// Package mdb Editors реализуются пока как студенты для понимания сути программы
// потом переписать под admins из примера
// TODO потом переписать логику на админов из примера
package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
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
func (r *EditorsRepo) Create(ctx context.Context, editor models.Editor) error {
	_, err := r.db.InsertOne(ctx, editor)
	return err
}

// GetByCredentials - Получение записи из коллекции по определенным полям
func (r *EditorsRepo) GetByCredentials(ctx context.Context, email, password models.Editor) error {
	return nil
}

// Verify - Подтверждение учетной зарегистрированной учетной записи
func (r *EditorsRepo) Verify(ctx context.Context, code string) error {
	codeId, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx,
		bson.M{"verification.code": codeId},
		bson.M{"$set": bson.M{"verification.verified": true}})

	return err
}
