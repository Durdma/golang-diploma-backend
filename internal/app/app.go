package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sas/internal/config"
	"sas/internal/controller"
	"sas/internal/repository"
	"sas/internal/server"
	"sas/internal/service"
	"sas/pkg/cache"
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

// Run Запуск всего приложения
func Run(configPath string, envPath string) {
	cfg, err := config.Init(configPath, envPath)
	if err != nil {
		panic(err)
	}

	if err := logger.Init(); err != nil {
		panic(err)
	}

	logger.Infof("%+v", *cfg)

	mongoClient := mongodb.NewClient("mongodb://localhost:27017", "", "")

	db := mongoClient.Database(cfg.Mongo.DatabaseName)
	memCache := cache.NewMemoryCache()

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos, memCache)

	handlers := controller.NewHandler(services.Universities, services.Editors)

	srv := server.NewServer(cfg, handlers.Init())
	go func() {
		if err := srv.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
	}

}
