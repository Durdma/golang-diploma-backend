package university

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
)

type StructEntity struct {
	Name      string
	published bool
}

// Institute - Сущность для описания института в университете.
type Institute struct {
	ID          primitive.ObjectID // id записи в MongoDB.
	Name        string             // Название института.
	Description string             // Описание института.
	ImageURL    string             // Ссылка на основное изображение.
	Director    Person             // Глава института.
	CreatedAt   int64              // Дата создания записи об институте.
	UpdatedAt   int64              // Дата последнего обновления записи об институте.
	Published   bool               // Статус публикации записи об институте.
	CreatedBy   models.User        // Автор записи об институте.
	UpdatedBy   []models.User      // Редакторы записи об институте.
}

// Department - Сущность для описания кафедры института.
type Department struct {
	ID          primitive.ObjectID // id записи в MongoDB.
	Name        string             // Название кафедры.
	Description string             // Описание кафедры.
	Staff       []Person           // Преподавательский состав кафедры.
	ImageURL    string             // Ссылка на основное изображение.
	CreatedAt   int64              // Дата создания записи о кафедре.
	UpdatedAt   int64              // Дата последнего обновления записи о кафедре.
	Published   bool               // Статус публикации записи о кафедре.
	CreatedBy   models.User        // Автор записи о кафедре.
	UpdatedBy   []models.User      // Редакторы записи о кафедре.
}

// Person - Сущность для описания сотрудника университета (информационная запись).
type Person struct {
	ID         primitive.ObjectID // id записи в MongoDB.
	Name       string             // Имя сотрудника.
	Surname    string             // Фамилия сотрудника.
	Patronymic string             //	Отчество сотрудника.
	ImageURL   string             // Ссылка на фото сотрудника.
	Info       string             // Дополнительная информация о сотруднике.
}
