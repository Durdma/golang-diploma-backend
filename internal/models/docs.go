package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Docs struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	UniversityId    primitive.ObjectID `json:"university_id" bson:"university_id"`
	Header          string             `json:"header" bson:"header"`
	Code            string             `json:"code" bson:"code"`
	Magistrate      bool               `json:"magistrate" bson:"magistrate"`
	Enrollee        bool               `json:"enrollee" bson:"enrollee"`
	Description     string             `json:"description" bson:"description"`
	DocURL          string             `json:"doc_url" bson:"doc_url"`
	PublicationDate time.Time          `json:"publication_date" bson:"publication_date"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

	CreatedBy primitive.ObjectID `json:"created_by" bson:"created_by"`
	UpdatedBy primitive.ObjectID `json:"updated_by" bson:"updated_by"`
}
