package admin

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
)

// Admin - сущность для сотрудников платформы
type Admin struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	models.BaseUser
}
