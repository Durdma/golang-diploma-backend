package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BaseUser - Основная сущность для всех учетных записей.
//type BaseUser struct {
//	Username string `json:"username" bson:"username"` // Имя пользователя
//	Password string `json:"password" bson:"password"` // Пароль
//	Email    string `json:"email" bson:"email"`       // Электронная почта
//}

// User - Сущность для хранения в БД.
type User struct {
	ID           primitive.ObjectID  `json:"id" bson:"_id"` // id записи в MongoBD.
	Name         string              `json:"name" bson:"name"`
	Email        string              `json:"email" bson:"email"`
	Password     string              `json:"hash" bson:"hash"`
	RegisteredAt primitive.Timestamp `json:"registered_at" bson:"registered_at"` // Дата регистрации
	Universities []University        `json:"universities" bson:"universities"`
	//BaseUser
}
