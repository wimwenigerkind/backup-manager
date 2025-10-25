package main

import (
	"fmt"
	"log"

	"github.com/wimwenigerkind/backup-manager/agent/internal/client"
	"github.com/wimwenigerkind/backup-manager/agent/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	printStart(cfg)
	apiClient := client.NewAPIClient(cfg.ServerURL, cfg.AgentID)
	configManager := config.NewManager(apiClient)
}

func printStart(cfg *config.Config) {
	log.Println("Starting Agent...")
	fmt.Println()
	log.Printf("Server URL: %s", cfg.ServerURL)
	log.Printf("Agent ID: %s", cfg.AgentID)
	log.Printf("Poll Interval: %v", cfg.PollInterval)
	fmt.Println()
}
