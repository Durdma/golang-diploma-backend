package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sas/internal/config"
	"sas/internal/controller/httpv1"
	"sas/internal/server"
	"sas/pkg/database/mongodb"
	"sas/pkg/logger"
	"syscall"
)

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

const configPath = "..\\..\\configs\\main"
const envPath = "../../app"

func main() {
	cfg, err := config.Init(configPath, envPath)
	if err != nil {
		panic(err)
	}

	if err := logger.Init(); err != nil {
		panic(err)
	}

	mongoClient := mongodb.NewClient("mongodb://localhost:27017", "", "")

	_ = mongoClient.Database("universityPlatform")

	handlers := httpv1.NewHandler()

	srv := server.NewServer(cfg, handlers.Init())
	go func() {
		if err := srv.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started!")

	// Конструкция для безопасного завершения работы сервиса. Отработает в любом случае.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
	}
}
