package university

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
)

// University - Сущность для университетов, расположенных на платформе.
type University struct {
	ID           primitive.ObjectID `json:"id" bson:"id"`                       // id записи в MongoDB.
	Name         string             `json:"name" bson:"name"`                   // Название университета.
	Description  string             `json:"description" bson:"description"`     // Описание университета.
	Domain       string             `json:"domain" bson:"domain"`               // Доменное имя университета.
	Verified     bool               `json:"verified" bson:"verified"`           // Статус верификации университета на платформе.
	RegisteredAt int64              `json:"registered_at" bson:"registered_at"` // Дата регистрации университета на платформе.
	VerifiedAt   int64              `json:"verified_at" bson:"verified_at"`     // Дата верификации университета на платформе.
	Editors      []Editor           `json:"editors" bson:"editors"`             // Список редакторов контента на сайте университета.
}

// Editor - Сущность редактора контента на сайте университета.
type Editor struct {
	ID primitive.ObjectID `json:"id" bson:"id"` // id записи в MongoDB.
	models.BaseUser
}
