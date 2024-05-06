package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Domain struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	HTTPDomainName string             `bson:"http_domain_name"`
	DBDomainName   string             `bson:"db_domain_name"`
	Visible        bool               `bson:"visible"`
	Deleted        bool               `bson:"deleted"`
}
