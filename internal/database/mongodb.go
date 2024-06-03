package database

import (
	"context"
	"log"
	"mc-insta-bot/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {
	cfg := config.LoadMongoConfig()
	client := createMongoClient(cfg)
	pingMongoClient(client, cfg)
	MongoClient = client
	log.Println("Connected to MongoDB!")
}

func createMongoClient(cfg config.MongoConfig) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoConnectTimeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	return client
}

func pingMongoClient(client *mongo.Client, cfg config.MongoConfig) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoPingTimeout)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
}
