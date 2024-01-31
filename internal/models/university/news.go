package university

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsEntity struct {
	Header      string `json:"header" bson:"header"`
	Description string `json:"description" bson:"description"`
	Published   bool   `json:"published" bson:"published"`
}

// News - Сущность для новостных записей университета.
type News struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id"`                  // id записи в MongoDB.
	Header      string              `json:"header" bson:"header"`           // Заголовок новостной записи.
	Description string              `json:"description" bson:"description"` // Краткое описание новостной записи.
	Body        string              `json:"body" bson:"body"`               // Основной текст новостной записи.
	ImageURL    string              `json:"image_url" bson:"image_url"`     // Ссылка на основное изображение новостной записи.
	CreatedAt   primitive.Timestamp `json:"created_at" bson:"created_at"`   // Дата создания новостной записи.
	UpdatedAt   primitive.Timestamp `json:"updated_at" bson:"updated_at"`   // Дата последнего обновления новостной записи.
	Published   bool                `json:"published" bson:"published"`     // Статус публикации новостной записи.
	CreatedBy   Editor              `json:"created_by" bson:"created_by"`   // Автор новостной записи.
	UpdatedBy   []Editor            `json:"updated_by" bson:"updated_by"`   // Редакторы новостной записи.
}
