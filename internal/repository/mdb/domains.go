package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sas/internal/models"
)

type DomainsRepo struct {
	db *mongo.Collection
}

func NewDNSRepo(db *mongo.Database) *DomainsRepo {
	return &DomainsRepo{
		db: db.Collection(domainsCollection),
	}
}

func (r *DomainsRepo) Create(ctx context.Context, domain models.Domain) error {
	_, err := r.db.InsertOne(ctx, domain)

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

	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &domains); err != nil {
		return nil, err
	}

	return domains, err
}
