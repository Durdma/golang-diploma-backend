package university

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseEntity struct {
	Name      string
	Code      string
	Published bool
}

// Course - Сущность для описания образовательной программы.
// Содержит полную информацию и ссылки на необходимые документы для описания образовательной программы.
type Course struct {
	ID           primitive.ObjectID // id в MongoDB.
	Name         string             // Название образовательной программы.
	Code         string             // Код образовательной программы.
	Description  string             // Описание образовательной программы
	Degree       Degree             // Уровень образовательной программы.
	DocumentsURL []string           // Ссылки на документы, необходимые для описания образовательной программы.
	ImageURL     string             // Ссылка на изображение для образовательной программы.
	CreatedAt    int64              // Дата создания записи.
	UpdatedAt    int64              // Дата изменения записи.
	Published    bool               // Статус публикации записи.
	CreatedBy    Editor             // Автор записи.
	UpdatedBy    []Editor           // Редакторы записи.
}

// Degree - Уровень образовательной программы.
type Degree struct {
	ID   primitive.ObjectID
	Name string
}
