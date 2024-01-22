package university

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
)

// University - Сущность для университетов, расположенных на платформе.
type University struct {
	ID           primitive.ObjectID // id записи в MongoDB.
	Name         string             // Название университета.
	Description  string             // Описание университета.
	Domain       string             // Доменное имя университета.
	Verified     bool               // Статус верификации университета на платформе.
	RegisteredAt int64              // Дата регистрации университета на платформе.
	VerifiedAt   int64              // Дата верификации университета на платформе.
	Editors      []Editor           // Список редакторов контента на сайте университета.
}

// Editor - Сущность редактора контента на сайте университета.
type Editor struct {
	ID primitive.ObjectID // id записи в MongoDB.
	models.BaseUser
}
