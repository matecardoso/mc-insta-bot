package main

import (
	"fmt"
	"log"
	"mc-insta-bot/internal/app"
	"mc-insta-bot/internal/config"
	"mc-insta-bot/internal/database"
)

func main() {
	log.Println("Starting application...")

	appConfig := config.LoadAppConfig()
	log.Printf("Loaded app config: %+v\n", appConfig)

	mongoConfig := config.LoadMongoConfig()
	log.Printf("Loaded mongo config: %+v\n", mongoConfig)

	database.ConnectMongo()
	log.Println("Connected to MongoDB")

	app := app.Setup()
	port := appConfig.ServerPort
	log.Printf("Starting server on port %s\n", port)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
