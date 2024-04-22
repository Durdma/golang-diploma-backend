package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type DNSRepo struct {
	db *mongo.Collection
}

func NewDNSRepo(db *mongo.Database) *DNSRepo {
	return &DNSRepo{
		db: db.Collection(domainsCollection),
	}
}

func (r *DNSRepo) Create(ctx context.Context, domain string) error {
	dbDomain := strings.ReplaceAll(domain, "-", "_")

	_, err := r.db.InsertOne(ctx, bson.D{
		{"domain_name", domain},
		{"db_domain_name", dbDomain},
	})
	return err
}

func (r *DNSRepo) Delete(ctx context.Context, domain string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{
		"domain_name": domain,
	})

	return err
}

func (r *DNSRepo) Get(ctx context.Context, domain string) (string, error) {
	domainName := make(map[string]string)
	domainName["db_domain_name"] = ""
	opts := options.FindOne().SetProjection(bson.D{
		{"domain_name", 0},
		{"_id", 0},
	})

	err := r.db.FindOne(ctx, bson.M{
		"domain_name": domain,
	}, opts).Decode(&domainName)

	return domainName["db_domain_name"], err
}
