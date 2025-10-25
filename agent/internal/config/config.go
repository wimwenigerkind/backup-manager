package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerURL    string
	AgentID      string
	PollInterval time.Duration
}

func LoadConfig() *Config {
	if _, err := os.Stat(".env"); err == nil {
		log.Println("Loading .env file")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return &Config{
		ServerURL:    getEnv("SERVER_URL", "http://localhost:8080"),
		AgentID:      getEnv("AGENT_ID", ""),
		PollInterval: time.Duration(getEnvInt("POLL_INTERVAL", 15)) * time.Second,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var intValue int
		_, err := fmt.Sscanf(value, "%d", &intValue)
		if err == nil {
			return intValue
		}
	}
	return defaultValue
}
