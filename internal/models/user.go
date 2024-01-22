package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// BaseUser - Основная сущность для всех учетных записей.
type BaseUser struct {
	Username string // Имя пользователя
	Password string // Пароль
	Email    string // Электронная почта
}

// User - Сущность для хранения в БД.
type User struct {
	ID           primitive.ObjectID // id записи в MongoBD.
	RegisteredAt int64              // Дата регистрации
	BaseUser
}
