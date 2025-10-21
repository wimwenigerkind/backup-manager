package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wimwenigerkind/backup-manager/server/internal/config"
	"github.com/wimwenigerkind/backup-manager/server/internal/database"
	"github.com/wimwenigerkind/backup-manager/server/internal/handlers"
)

func main() {
	cfg := config.LoadConfig()

	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	agentHandler := handlers.NewAgentHandler()
	api := r.Group("/api/v1")
	{
		agents := api.Group("/agents")
		{
			agents.POST("", agentHandler.CreateAgent)
			agents.GET("", agentHandler.GetAgents)
		}
	}

	if err := r.Run(); err != nil {
		log.Fatal("Failed to start gin:", err)
	}
}
