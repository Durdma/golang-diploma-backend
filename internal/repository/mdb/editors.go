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
		db: db.Collection(usersCollection),
	}
}

// Create - Добавление записи о новом редакторе в коллекцию
func (r *EditorsRepo) Create(ctx context.Context, editor models.Editor) error {
	_, err := r.db.InsertOne(ctx, editor)
	return err
}

func (r *EditorsRepo) ChangeBlockStatus(ctx context.Context, editorId string, state bool) error {
	id, err := primitive.ObjectIDFromHex(editorId)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx, bson.M{
		"_id": id,
	},
		bson.M{
			"$set": bson.M{"is_blocked": state},
		})

	return err
}

func (r *EditorsRepo) ChangeVerificationStatus(ctx context.Context, editorId string, state bool) error {
	id, err := primitive.ObjectIDFromHex(editorId)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx, bson.M{
		"_id": id,
	},
		bson.M{
			"$set": bson.M{"verification.verified": state},
		})

	return err
}

func (r *EditorsRepo) GetEditorById(ctx context.Context, userId primitive.ObjectID) (models.User, error) {
	var user models.User
	err := r.db.FindOne(ctx, bson.M{
		"_id":      userId,
		"is_admin": false,
	}).Decode(&user)

	return user, err
}

func (r *EditorsRepo) UpdateEditor(ctx context.Context, user models.User) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": user.ID,
	}, bson.M{"$set": bson.M{
		"name":                  user.Name,
		"email":                 user.Email,
		"password":              user.Password,
		"domain_id":             user.DomainId,
		"is_blocked":            user.IsBlocked,
		"verification.verified": user.Verification.Verified,
	}})

	return err
}

func (r *EditorsRepo) GetAllEditors(ctx context.Context) ([]models.User, error) {
	var editors []models.User

	filter := bson.M{
		"is_admin": false,
	}

	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &editors); err != nil {
		return nil, err
	}

	return editors, err
}
