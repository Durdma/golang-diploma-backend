package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
	"time"
)

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		db: db.Collection(usersCollection),
	}
}

func (r *UsersRepo) Create(ctx context.Context, user models.User) error {
	_, err := r.db.InsertOne(ctx, user)
	return err
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, email string, password string, domain primitive.ObjectID) (models.User, error) {
	var user models.User
	err := r.db.FindOne(ctx, bson.M{
		"email":                 email,
		"password":              password,
		"verification.verified": true,
		"is_blocked":            false,
		"domain_id":             domain,
	}).Decode(&user)

	return user, err
}

func (r *UsersRepo) GetByRefreshToken(ctx context.Context, domain primitive.ObjectID, refreshToken string) (models.User, error) {
	var user models.User
	err := r.db.FindOne(ctx, bson.M{
		"domain":                domain,
		"session.refresh_token": refreshToken,
		"session.expires_at": bson.M{
			"$gt": time.Now(),
		},
	}).Decode(&user)

	return user, err
}

func (r *UsersRepo) SetSession(ctx context.Context, userId primitive.ObjectID, session models.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": userId}, bson.M{"$set": bson.M{"session": session}})

	return err
}

func (r *UsersRepo) Verify(ctx context.Context, domain primitive.ObjectID, code string) error {
	codeId, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx,
		bson.M{
			"domain":            domain,
			"verification.code": codeId,
		},
		bson.M{"$set": bson.M{"verification.verified": true}})

	return err
}

func (r *UsersRepo) GetUserById(ctx context.Context, userId primitive.ObjectID) (models.User, error) {
	var user models.User
	err := r.db.FindOne(ctx, bson.M{
		"_id": userId,
	}).Decode(&user)

	return user, err
}

func (r *UsersRepo) GetAllEditors(ctx context.Context) ([]models.User, error) {
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
