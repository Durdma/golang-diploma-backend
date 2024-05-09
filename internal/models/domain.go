package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//type Domain struct {
//	ID             primitive.ObjectID `bson:"_id,omitempty"`
//	HTTPDomainName string             `bson:"http_domain_name"`
//	DBDomainName   string             `bson:"db_domain_name"`
//	Visible        bool               `bson:"visible"`
//	Deleted        bool               `bson:"deleted"`
//}

type Domain struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	SiteId         primitive.ObjectID `json:"-" bson:"site_id,omitempty"`
	HTTPDomainName string             `json:"http_domain_name" bson:"http_domain_name"`
	DBDomainName   string             `json:"db_domain_name" bson:"db_domain_name"`
	DomainName     string             `json:"domain_name" bson:"domain_name"`
	ShortName      string             `json:"short_name" bson:"short_name"`
	Visible        bool               `json:"visible" bson:"visible"`
	Deleted        bool               `json:"deleted" bson:"deleted"`
	Verified       bool               `json:"verified" bson:"verified"`
}
