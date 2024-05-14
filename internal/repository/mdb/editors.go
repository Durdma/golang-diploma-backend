// Package mdb Editors реализуются пока как студенты для понимания сути программы
// потом переписать под admins из примера
// TODO потом переписать логику на админов из примера
package mdb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sas/internal/models"
	"strconv"
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
func (r *EditorsRepo) Create(ctx context.Context, editor models.User) error {
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
	var sort bson.D

	aggregateState := false

	filter := bson.M{
		"is_admin": false,
	}

	if val := ctx.Value("name"); val != nil {
		pattern := fmt.Sprintf(".*%s.*", val.(string))
		filter["name"] = bson.M{"$regex": pattern, "$options": "i"}
	}

	if val := ctx.Value("university"); val != nil {
		if val.(string) != "any" {
			universityId, err := primitive.ObjectIDFromHex(val.(string))
			if err != nil {
				return nil, err
			}
			filter["domain_id"] = universityId
		}
	}

	if val := ctx.Value("verify"); val != nil {
		if val.(string) == "true" || val.(string) == "false" {
			verified, err := strconv.ParseBool(val.(string))
			if err != nil {
				return nil, err
			}
			fmt.Println("verification ", verified)
			filter["verification.verified"] = verified
		}
	}

	if val := ctx.Value("block"); val != nil {
		if val.(string) == "true" || val.(string) == "false" {
			blocked, err := strconv.ParseBool(val.(string))
			if err != nil {
				return nil, err
			}
			fmt.Println("block ", blocked)
			filter["is_blocked"] = blocked
		}
	}

	if val := ctx.Value("sort"); val != nil {
		switch val.(string) {
		case "name":
			sort = bson.D{{"name", 1}}
		case "university":
			aggregateState = true
		case "entry":
			sort = bson.D{{"last_visited_at", 1}}
		case "registration":
			sort = bson.D{{"registered_at", 1}}
		}
	} else {
		sort = bson.D{{"name", 1}}
	}

	if !aggregateState {
		opts := options.Find().SetSort(sort)
		cursor, err := r.db.Find(ctx, filter, opts)
		if err != nil {
			return nil, err
		}

		if err = cursor.All(ctx, &editors); err != nil {
			return nil, err
		}

		return editors, err
	}
	matchStage := bson.D{{"$match", filter}}

	lookUpStage := bson.D{
		{"$lookup", bson.D{
			{"from", "domains"},
			{"localField", "domain_id"},
			{"foreignField", "_id"},
			{"as", "domain"},
		}}}

	// unWindStage := bson.D{{"$unwind", bson.D{{"path"}}}}

	sortStage := bson.D{{"$sort", bson.D{{"domain.domain_name", 1}}}}

	cursor, err := r.db.Aggregate(ctx, mongo.Pipeline{matchStage, lookUpStage, sortStage})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &editors); err != nil {
		return nil, err
	}

	return editors, nil
}
