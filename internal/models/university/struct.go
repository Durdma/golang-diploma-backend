package university

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StructEntity struct {
	Name      string
	published bool
}

// Institute - Сущность для описания института в университете.
type Institute struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id"`                  // id записи в MongoDB.
	Name        string              `json:"name" bson:"name"`               // Название института.
	Description string              `json:"description" bson:"description"` // Описание института.
	ImageURL    string              `json:"image_url" bson:"image_url"`     // Ссылка на основное изображение.
	Director    Person              `json:"director" bson:"director"`       // Глава института.
	CreatedAt   primitive.Timestamp `json:"created_at" bson:"created_at"`   // Дата создания записи об институте.
	UpdatedAt   primitive.Timestamp `json:"updated_at" bson:"updated_at"`   // Дата последнего обновления записи об институте.
	Published   bool                `json:"published" bson:"published"`     // Статус публикации записи об институте.
	CreatedBy   Editor              `json:"created_by" bson:"created_by"`   // Автор записи об институте.
	UpdatedBy   []Editor            `json:"updated_by" bson:"updated_by"`   // Редакторы записи об институте.
}

// Department - Сущность для описания кафедры института.
type Department struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id"`                  // id записи в MongoDB.
	Name        string              `json:"name" bson:"name"`               // Название кафедры.
	Description string              `json:"description" bson:"description"` // Описание кафедры.
	Staff       []Person            `json:"staff" bson:"staff"`             // Преподавательский состав кафедры.
	ImageURL    string              `json:"image_url" bson:"image_url"`     // Ссылка на основное изображение.
	CreatedAt   primitive.Timestamp `json:"created_at" bson:"created_at"`   // Дата создания записи о кафедре.
	UpdatedAt   primitive.Timestamp `json:"updated_at" bson:"updated_at"`   // Дата последнего обновления записи о кафедре.
	Published   bool                `json:"published" bson:"published"`     // Статус публикации записи о кафедре.
	CreatedBy   Editor              `json:"created_by" bson:"created_by"`   // Автор записи о кафедре.
	UpdatedBy   []Editor            `json:"updated_by" bson:"updated_by"`   // Редакторы записи о кафедре.
}

// Person - Сущность для описания сотрудника университета (информационная запись).
type Person struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`                // id записи в MongoDB.
	Name       string             `json:"name" bson:"name"`             // Имя сотрудника.
	Surname    string             `json:"surname" bson:"surname"`       // Фамилия сотрудника.
	Patronymic string             `json:"patronymic" bson:"patronymic"` //	Отчество сотрудника.
	ImageURL   string             `json:"image_url" bson:"image_url"`   // Ссылка на фото сотрудника.
	Info       string             `json:"info" bson:"info"`             // Дополнительная информация о сотруднике.
}
