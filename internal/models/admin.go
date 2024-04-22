package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Admin - сущность для сотрудников платформы
type Admin struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	RegisteredAt time.Time          `json:"registered_at" bson:"registered_at"`
	LastVisitAt  time.Time          `json:"last_visit_at" bson:"last_visit_at"`
	Verification Verification       `json:"verification" bson:"verification"`
	Session      Session            `json:"session" bson:"session"`
	//models.BaseUser
}
