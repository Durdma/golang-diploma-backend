package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
	"sas/pkg/logger"
	"time"
)

type AdminsRepo struct {
	db *mongo.Collection
}

func NewAdminsRepo(db *mongo.Database) *AdminsRepo {
	return &AdminsRepo{
		db: db.Collection(adminsCollection),
	}
}

func (r *AdminsRepo) Create(ctx context.Context, adm models.Admin) error {
	_, err := r.db.InsertOne(ctx, adm)
	return err
}

func (r *AdminsRepo) GetByCredentials(ctx context.Context, email string, password string) (models.Admin, error) {
	var adm models.Admin
	err := r.db.FindOne(ctx, bson.M{
		"email":        email,
		"password":     password,
		"verification": true,
	}).Decode(&adm)

	logger.Infof("%+v\n", adm.ID)

	return adm, err
}

func (r *AdminsRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (models.Admin, error) {
	var adm models.Admin
	err := r.db.FindOne(ctx, bson.M{
		"session.refresh_token": refreshToken,
		"session.expires_at": bson.M{
			"$gt": time.Now(),
		},
	}).Decode(&adm)

	return adm, err
}

func (r *AdminsRepo) SetSession(ctx context.Context, userId primitive.ObjectID, session models.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": userId}, bson.M{"$set": bson.M{"session": session}})
	logger.Info(userId)

	return err
}

func (r *AdminsRepo) Verify(ctx context.Context, code string) error {
	codeId, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx,
		bson.M{"verification.code": codeId},
		bson.M{"$set": bson.M{"verification.verified": true}})

	return err
}
