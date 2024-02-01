package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const timeout = 10 * time.Second // Максимальное Время обработки запроса к БД

// NewClient - Создание нового подключения к БД
func NewClient(uri string, username string, password string) *mongo.Client {
	opts := options.Client().ApplyURI(uri) // Добавление значений для подключения к БД

	// Установка максимального времени обработки запроса
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Подключение к БД
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Error while connecting to MongoDB!")
	}

	// Тестовый запрос к БД
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return client
}
