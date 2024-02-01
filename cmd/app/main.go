package main

import "sas/internal/app"

// @title University Platform API
// @version 1.0
// @description API Server for University Platform

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey AdminAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey EditorsAuth
// @in header
// @name Authorization

const configPath = "..\\..\\configs\\main" // configPath - путь к конфигу параметров API
const envPath = "../../app"                // envPath - путь к переменным окружения, где хранится чувствительная информация. НЕ ЗАЛИВАТЬ НА ГИТ!!!

// main - запуск всего приложения
func main() {
	app.Run(configPath, envPath)
}
