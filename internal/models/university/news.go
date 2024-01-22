package university

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
)

type NewsEntity struct {
	Header      string
	Description string
	Published   bool
}

// News - Сущность для новостных записей университета.
type News struct {
	ID          primitive.ObjectID // id записи в MongoDB.
	Header      string             // Заголовок новостной записи.
	Description string             // Краткое описание новостной записи.
	Body        string             // Основной текст новостной записи.
	ImageURL    string             // Ссылка на основное изображение новостной записи.
	CreatedAt   int64              // Дата создания новостной записи.
	UpdatedAt   int64              // Дата последнего обновления новостной записи.
	Published   bool               // Статус публикации новостной записи.
	CreatedBy   models.User        // Автор новостной записи.
	UpdatedBy   []models.User      // Редакторы новостной записи.
}
