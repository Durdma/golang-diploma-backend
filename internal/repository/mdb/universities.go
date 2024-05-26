package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
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

func (r *UniversityRepo) Create(ctx context.Context, university models.University) (primitive.ObjectID, error) {
	res, err := r.db.InsertOne(ctx, university)

	return res.InsertedID.(primitive.ObjectID), err
}

// GetByDomain - Получение записи об университете по имени домена
func (r *UniversityRepo) GetByDomain(ctx context.Context, domainName string) (models.University, error) {
	var univ models.University
	err := r.db.FindOne(ctx, bson.M{
		"domain": domainName,
	}).Decode(&univ)

	return univ, err
}

func (r *UniversityRepo) GetByUniversityId(ctx context.Context, siteId primitive.ObjectID) (models.University, error) {
	var university models.University
	err := r.db.FindOne(ctx, bson.M{
		"_id": siteId,
	}).Decode(&university)

	return university, err
}

func (r *UniversityRepo) UpdateUniversity(ctx context.Context) error {
	return nil
}

func (r *UniversityRepo) ChangeCSS(ctx context.Context, universityId primitive.ObjectID, colors map[string]string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": universityId,
	},
		bson.M{
			"$set": bson.M{
				"settings.main_color":                   colors["main_color"],
				"settings.main_color_hover":             colors["main_color_hover"],
				"settings.main_footer_font_color":       colors["main_footer_font_color"],
				"settings.main_footer_font_color_hover": colors["main_footer_font_color_hover"],
				"settings.main_footer_bg_color":         colors["main_footer_bg_color"],
			},
		},
	)

	return err
}

func (r *UniversityRepo) SetUniversityHistory(ctx context.Context, universityId primitive.ObjectID, history models.History) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": universityId,
	},
		bson.M{
			"$set": bson.M{
				"history": history,
			},
		},
	)

	return err
}
