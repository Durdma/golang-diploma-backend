package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Domain struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SiteId         primitive.ObjectID `json:"-" bson:"site_id,omitempty"`
	HTTPDomainName string             `json:"http_domain_name" bson:"http_domain_name"`
	DBDomainName   string             `json:"db_domain_name" bson:"db_domain_name"`
	DomainName     string             `json:"domain_name" bson:"domain_name"`
	ShortName      string             `json:"short_name" bson:"short_name"`
	Visible        bool               `json:"visible" bson:"visible"`
	Verified       bool               `json:"verified" bson:"verified"`
	RegisteredAt   time.Time          `json:"registered_at" bson:"registered_at"`
	LastUpdate     time.Time          `json:"last_update" bson:"last_update"`
}
