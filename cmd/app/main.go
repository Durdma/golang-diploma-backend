package main

import (
	"context"
	"fmt"
	"sas/internal/config"
	"sas/pkg/database/mongodb"
)

// TODO смотри в config.go
const configPath = "D:\\Projects\\University\\LAST_COURSE\\ДП\\go-saas\\configs\\main.yml"

func main() {
	mongoClient := mongodb.NewClient("mongodb://localhost:27017", "", "")
	defer mongoClient.Disconnect(context.Background())

	db := mongoClient.Database("universityPlatform")

	cfg, err := config.Init(configPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)

	fmt.Println(db.CreateCollection(context.Background(), "admins"))

}
