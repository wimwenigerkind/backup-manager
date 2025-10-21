package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wimwenigerkind/backup-manager/server/internal/config"
	"github.com/wimwenigerkind/backup-manager/server/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	db(cfg)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatal("Failed to start gin:", err)
	}
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
