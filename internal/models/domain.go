package models

type Domain struct {
	HTTPDomainName string `bson:"http_domain_name"`
	DBDomainName   string `bson:"db_domain_name"`
	Visible        bool   `bson:"visible"`
	Deleted        bool   `bson:"deleted"`
}
