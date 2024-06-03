package config

import (
	"encoding/json"
	"io"
	"log"
	"mc-insta-bot/internal/utils"
	"os"
)

type AppConfig struct {
	ServerPort string `json:"server_port"`
}

func LoadAppConfig() AppConfig {
	file, err := os.Open("configs/app.json")
	if err != nil {
		log.Fatalf("Error reading app config file: %v", err)
	}
	defer file.Close()

	configData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading app config data: %v", err)
	}

	configDataStr := utils.ReplacePlaceholders(string(configData))

	var config AppConfig
	if err := json.Unmarshal([]byte(configDataStr), &config); err != nil {
		log.Fatalf("Error parsing app config file: %v", err)
	}

	return config
}
