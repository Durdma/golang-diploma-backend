package admin

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Admin - сущность для сотрудников платформы
type Admin struct {
	ID           primitive.ObjectID  `json:"id" bson:"_id"`
	Name         string              `json:"name" bson:"name"`
	Email        string              `json:"email" bson:"email"`
	Password     string              `json:"password" bson:"password"`
	RegisteredAt primitive.Timestamp `json:"registered_at" bson:"registered_at"`
	LastVisitAt  primitive.Timestamp `json:"last_visit_at" bson:"last_visit_at"`
	//models.BaseUser
}
