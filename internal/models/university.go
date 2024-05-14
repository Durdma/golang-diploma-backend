package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// University - Сущность для университетов, расположенных на платформе.
type University struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"` // id записи в MongoDB.
	Name        string             `json:"name" bson:"name"`        // Название университета.
	ShortName   string             `json:"short_name" bson:"short_name"`
	Description string             `json:"description" bson:"description"` // Описание университета.
	//Domain       Domain              `json:"domain" bson:"domain"`               // Доменное имя университета.
	Verification bool      `json:"verification" bson:"verification"`   // Статус верификации университета на платформе.
	RegisteredAt time.Time `json:"registered_at" bson:"registered_at"` // Дата регистрации университета на платформе.
	VerifiedAt   time.Time `json:"verified_at" bson:"verified_at"`     // Дата верификации университета на платформе.
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
	Editors      []Editor  `json:"editors" bson:"editors"` // Список редакторов контента на сайте университета.
	News         []News    `json:"news" bson:"news"`
}

// Editor - Сущность редактора контента на сайте университета.
type Editor struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"` // id записи в MongoDB.
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Verification Verification       `json:"verification" bson:"verification"` // Статус верификации
	RegisteredAt time.Time          `json:"registered_at" bson:"registered_at"`
	LastVisitAt  time.Time          `json:"last_visit_at" bson:"last_visit_at"`
	UniversityID primitive.ObjectID `json:"university_id" bson:"university_id"`
	Session      Session            `json:"session" bson:"session"`
}
