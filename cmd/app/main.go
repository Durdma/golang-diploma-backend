package main

import (
	"context"
	"fmt"
	"sas/pkg/database/mongodb"
)

func main() {
	mongoClient := mongodb.NewClient("mongodb://localhost:27017", "", "")
	defer mongoClient.Disconnect(context.Background())

	db := mongoClient.Database("universityPlatform")

	fmt.Println(db.CreateCollection(context.Background(), "admins"))
}
