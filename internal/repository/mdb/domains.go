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
	"time"
)

type DomainsRepo struct {
	db *mongo.Collection
}

// TODO rename to domains repo
func NewDNSRepo(db *mongo.Database) *DomainsRepo {
	return &DomainsRepo{
		db: db.Collection(domainsCollection),
	}
}

func (r *DomainsRepo) Create(ctx context.Context, domain models.Domain) (primitive.ObjectID, error) {
	res, err := r.db.InsertOne(ctx, domain)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), err
}

func (r *DomainsRepo) AddUniversityId(ctx context.Context, domainId primitive.ObjectID, universityId primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": domainId,
	},
		bson.M{
			"$set": bson.M{"site_id": universityId},
		})

	return err
}

// TODO refactor change status to 1 func
func (r *DomainsRepo) ChangeVisibleStatus(ctx context.Context, domainId string, state bool) error {
	id, err := primitive.ObjectIDFromHex(domainId)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx, bson.M{
		"_id": id,
	},
		bson.M{
			"$set": bson.M{"visible": state},
		})

	return err
}

func (r *DomainsRepo) ChangeVerificationStatus(ctx context.Context, domainId string, state bool) error {
	id, err := primitive.ObjectIDFromHex(domainId)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx, bson.M{
		"_id": id,
	},
		bson.M{
			"$set": bson.M{"verified": state},
		})

	return err
}

func (r *DomainsRepo) Delete(ctx context.Context, domain primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{
		"_id": domain,
	})
	if err != nil {
		return err
	}

	return err
}

func (r *DomainsRepo) GetByHTTPName(ctx context.Context, domain string) (models.Domain, error) {
	var resp models.Domain

	err := r.db.FindOne(ctx, bson.M{
		"http_domain_name": domain,
	}).Decode(&resp)
	if err != nil {
		return models.Domain{}, err
	}

	return resp, err
}

func (r *DomainsRepo) GetById(ctx context.Context, domainId primitive.ObjectID) (models.Domain, error) {
	var resp models.Domain

	err := r.db.FindOne(ctx, bson.M{
		"_id": domainId,
	}).Decode(&resp)
	if err != nil {
		return models.Domain{}, err
	}

	return resp, nil
}

func (r *DomainsRepo) GetByDomainName(ctx context.Context, domainName string) (models.Domain, error) {
	var domain models.Domain
	err := r.db.FindOne(ctx, bson.M{
		"domain_name": domainName,
	}).Decode(&domain)
	if err != nil {
		return models.Domain{}, err
	}

	return domain, err
}

func (r *DomainsRepo) GetAllDomains(ctx context.Context) ([]models.Domain, error) {
	var domains []models.Domain
	var sort bson.D

	filter := bson.M{}

	if val := ctx.Value("domain_name"); val != nil {
		pattern := fmt.Sprintf(".*%s.*", val.(string))
		filter["$or"] = []bson.M{
			{"domain_name": bson.M{"$regex": pattern, "$options": "i"}},
			{"short_name": bson.M{"$regex": pattern, "$options": "i"}},
		}
	}

	if val := ctx.Value("verify"); val != nil {
		if val.(string) == "true" || val.(string) == "false" {
			verified, err := strconv.ParseBool(val.(string))
			if err != nil {
				return nil, err
			}
			filter["verified"] = verified
		}
	}

	if val := ctx.Value("visible"); val != nil {
		if val.(string) == "true" || val.(string) == "false" {
			visible, err := strconv.ParseBool(val.(string))
			if err != nil {
				return nil, err
			}
			filter["visible"] = visible
		}
	}

	if val := ctx.Value("sort"); val != nil {
		switch val.(string) {
		case "name":
			sort = bson.D{{"domain_name", 1}}
		case "short-name":
			sort = bson.D{{"short_name", 1}}
		case "registration":
			sort = bson.D{{"registered_at", 1}}
		case "update":
			sort = bson.D{{"last_update", 1}}
		}
	} else {
		sort = bson.D{{"domain_name", 1}}
	}

	opts := options.Find().SetSort(sort)
	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &domains); err != nil {
		return nil, err
	}

	return domains, err
}

func (r *DomainsRepo) UpdateDomain(ctx context.Context, domain models.Domain) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id": domain.ID,
	}, bson.M{"$set": bson.M{
		"domain_name": domain.DomainName,
		"short_name":  domain.ShortName,
		"visible":     domain.Visible,
		"verified":    domain.Verified,
		"last_update": time.Now(),
	}})

	return err
}
