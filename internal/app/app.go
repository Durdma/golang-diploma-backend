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
	"sas/pkg/auth"
	"sas/pkg/cache"
	"sas/pkg/database/mongodb"
	"sas/pkg/email"
	"sas/pkg/hash"
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

// Run Запуск всего приложения, подключение всех компонентов
func Run(configPath string, envPath string) {
	// Подтягивание конфигурации
	cfg, err := config.Init(configPath, envPath)
	if err != nil {
		logger.Error(err)
		return
	}

	// Вывод конфигурации приложения
	logger.Infof("%+v", *cfg)

	// Подключение к mongoDB
	mongoClient := mongodb.NewClient("mongodb://localhost:27017", "", "")
	db := mongoClient.Database(cfg.Mongo.DatabaseName)

	// Подключение кэша
	memCache := cache.NewMemoryCache(int64(cfg.CacheTTL))

	// Подключение хэша
	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	emailProvider := email.NewClient(cfg.Email.Email, cfg.Email.Password, cfg.Email.Provider,
		cfg.Email.Port)

	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)
		return
	}

	// Подключение репозиториев
	repos := repository.NewRepositories(db)
	// Подключение сервисов
	services := service.NewServices(repos, memCache, hasher,
		tokenManager, emailProvider, cfg.Auth.JWT.AccessTokenTTL, cfg.Auth.JWT.RefreshTokenTTL)

	// Добавление контроллера
	handlers := controller.NewHandler(services.Universities, services.Editors, services.Admins, tokenManager, services.Domains, services.Users, services.Sites)

	// Инициализация сервера и его запуск
	srv := server.NewServer(cfg, handlers.Init(cfg.HTTP.Host, cfg.HTTP.Port))
	go func() {
		if err := srv.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Выключение сервера
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	// Отключение от БД
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
	}

}
