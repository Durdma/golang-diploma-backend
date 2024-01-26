package main

import (
	"context"
	"sas/internal/config"
	"sas/pkg/database/mongodb"
	"sas/pkg/logger"
)

const configPath = "../../configs/main"
const envPath = "../../app"

func main() {
	mongoClient := mongodb.NewClient("mongodb://localhost:27017", "", "")
	defer mongoClient.Disconnect(context.Background())

	db := mongoClient.Database("universityPlatform")

	cfg, err := config.Init(configPath, envPath)
	if err != nil {
		panic(err)
	}

	if err := logger.Init(); err != nil {
		panic(err)
	}

	logger.Infof("%+v\n", cfg)

	logger.Info(db.CreateCollection(context.Background(), "admins"))

}
