package config

import (
	"encoding/json"
	"io"
	"log"
	"mc-insta-bot/internal/utils"
	"os"
	"time"
)

type MongoConfig struct {
	MongoURI               string        `json:"mongo_uri"`
	MongoConnectTimeoutStr string        `json:"mongo_connect_timeout"`
	MongoPingTimeoutStr    string        `json:"mongo_ping_timeout"`
	MongoConnectTimeout    time.Duration `json:"-"`
	MongoPingTimeout       time.Duration `json:"-"`
}

func LoadMongoConfig() MongoConfig {
	file, err := os.Open("configs/mongo.json")
	if err != nil {
		log.Fatalf("Error reading mongo config file: %v", err)
	}
	defer file.Close()

	configData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading mongo config data: %v", err)
	}

	configDataStr := utils.ReplacePlaceholders(string(configData))

	var config MongoConfig
	if err := json.Unmarshal([]byte(configDataStr), &config); err != nil {
		log.Fatalf("Error parsing mongo config file: %v", err)
	}

	config.MongoConnectTimeout, err = time.ParseDuration(config.MongoConnectTimeoutStr)
	if err != nil {
		log.Fatalf("Invalid duration for MongoConnectTimeout: %v", err)
	}

	config.MongoPingTimeout, err = time.ParseDuration(config.MongoPingTimeoutStr)
	if err != nil {
		log.Fatalf("Invalid duration for MongoPingTimeout: %v", err)
	}

	return config
}
