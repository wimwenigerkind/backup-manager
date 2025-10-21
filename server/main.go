package main

import (
	"log"

	"github.com/wimwenigerkind/backup-manager/server/internal/config"
	"github.com/wimwenigerkind/backup-manager/server/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	db(cfg)
}

func db(cfg *config.Config) {
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatal("Failed to close database connection:", err)
		}
	}()

	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}
}
