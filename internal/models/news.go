package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// News - Сущность для новостных записей университета.

type News struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // id записи в MongoDB.
	DomainId    primitive.ObjectID `json:"university_id" bson:"university_id"`
	Header      string             `json:"header" bson:"header"`           // Заголовок новостной записи.
	Description string             `json:"description" bson:"description"` // Краткое описание новостной записи.
	Body        string             `json:"body" bson:"body"`               // Основной текст новостной записи.
	ImageURL    string             `json:"image_url" bson:"image_url"`     // Ссылка на основное изображение новостной записи.
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`   // Дата создания новостной записи.
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`   // Дата последнего обновления новостной записи.
	Published   bool               `json:"published" bson:"published"`     // Статус публикации новостной записи.
	CreatedBy   primitive.ObjectID `json:"created_by" bson:"created_by"`   // Автор новостной записи.
	UpdatedBy   primitive.ObjectID `json:"updated_by" bson:"updated_by"`   // Редакторы новостной записи.
}
