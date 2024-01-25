package main

import (
	"context"
	"fmt"
	"sas/internal/config"
	"sas/pkg/database/mongodb"
)

const configPath = "..\\..\\configs\\main"
const envPath = "../../app"

func main() {
	mongoClient := mongodb.NewClient("mongodb://localhost:27017", "", "")
	defer mongoClient.Disconnect(context.Background())

	db := mongoClient.Database("universityPlatform")

	cfg, err := config.Init(configPath, envPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)

	fmt.Println(db.CreateCollection(context.Background(), "admins"))

}
