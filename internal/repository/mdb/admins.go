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

type AdminsRepo struct {
	db *mongo.Collection
}

func NewAdminsRepo(db *mongo.Database) *AdminsRepo {
	return &AdminsRepo{
		db: db.Collection(usersCollection),
	}
}

func (r *AdminsRepo) Create(ctx context.Context, adm models.User) error {
	_, err := r.db.InsertOne(ctx, adm)
	return err
}

func (r *AdminsRepo) GetAll(ctx context.Context) ([]models.User, error) {
	var admins []models.User
	var sort bson.D

	filter := bson.M{
		"is_admin": true,
	}

	if val := ctx.Value("name"); val != nil {
		pattern := fmt.Sprintf(".*%s.*", val.(string))
		filter["name"] = bson.M{"$regex": pattern, "$options": "i"}
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
		case "entry":
			sort = bson.D{{"last_visited_at", 1}}
		case "registration":
			sort = bson.D{{"registered_at", 1}}
		}
	} else {
		sort = bson.D{{"name", 1}}
	}

	opts := options.Find().SetSort(sort)
	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &admins); err != nil {
		return nil, err
	}

	return admins, err
}

func (r *AdminsRepo) GetAdminById(ctx context.Context, adminId primitive.ObjectID) (models.User, error) {
	var user models.User
	err := r.db.FindOne(ctx, bson.M{
		"_id":      adminId,
		"is_admin": true,
	}).Decode(&user)

	return user, err
}

func (r *AdminsRepo) ChangeBlockStatus(ctx context.Context, adminId string, state bool) error {
	id, err := primitive.ObjectIDFromHex(adminId)
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

func (r *AdminsRepo) ChangeVerificationStatus(ctx context.Context, editorId string, state bool) error {
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

func (r *AdminsRepo) UpdateAdmin(ctx context.Context, user models.User) error {
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
