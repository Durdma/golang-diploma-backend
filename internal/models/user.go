package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	DomainId     primitive.ObjectID `json:"domain_id" bson:"domain_id"`
	DomainName   string             `json:"domain_name" bson:"domain.domain_name,omitempty"`
	IsAdmin      bool               `json:"is_admin" bson:"is_admin"`
	IsBlocked    bool               `json:"is_blocked" bson:"is_blocked"`
	RegisteredAt time.Time          `json:"registered_at" bson:"registered_at"`
	LastVisitAt  time.Time          `json:"last_visit_at" bson:"last_visit_at"`
	Verification Verification       `json:"verification" bson:"verification"`
	Session      Session            `json:"session" bson:"session"`
}
