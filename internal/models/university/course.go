package university

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseEntity struct {
	Name      string `json:"name" bson:"name"`
	Code      string `json:"code" bson:"code"`
	Published bool   `json:"published" bson:"published"`
}

// Course - Сущность для описания образовательной программы.
// Содержит полную информацию и ссылки на необходимые документы для описания образовательной программы.
type Course struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`                      // id в MongoDB.
	Name         string             `json:"name" bson:"name"`                   // Название образовательной программы.
	Code         string             `json:"code" bson:"code"`                   // Код образовательной программы.
	Description  string             `json:"description" bson:"description"`     // Описание образовательной программы
	Degree       Degree             `json:"degree" bson:"degree"`               // Уровень образовательной программы.
	DocumentsURL []string           `json:"documents_url" bson:"documents_url"` // Ссылки на документы, необходимые для описания образовательной программы.
	ImageURL     string             `json:"image_url" bson:"image_url"`         // Ссылка на изображение для образовательной программы.
	CreatedAt    int64              `json:"created_at" bson:"created_at"`       // Дата создания записи.
	UpdatedAt    int64              `json:"updated_at" bson:"updated_at"`       // Дата изменения записи.
	Published    bool               `json:"published" bson:"published"`         // Статус публикации записи.
	CreatedBy    Editor             `json:"created_by" bson:"created_by"`       // Автор записи.
	UpdatedBy    []Editor           `json:"updated_by" bson:"updated_by"`       // Редакторы записи.
}

// Degree - Уровень образовательной программы.
type Degree struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Name string             `json:"name" bson:"name"`
}
