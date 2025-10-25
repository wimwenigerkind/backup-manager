package config

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/wimwenigerkind/backup-manager/agent/internal/client"
	"github.com/wimwenigerkind/backup-manager/agent/internal/models"
)

type Manager struct {
	client        *client.APIClient
	currentConfig *models.AgentConfigResponse
	configHash    string
	mu            sync.RWMutex
}

func NewManager(apiClient *client.APIClient) *Manager {
	return &Manager{
		client: apiClient,
	}
}

func (m *Manager) StartPolling(ctx context.Context, interval time.Duration, onUpdate func([]models.BackupJob)) {
	log.Printf("Starting config polling with interval %v", interval)

	m.fetchAndUpdate(onUpdate)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.fetchAndUpdate(onUpdate)
		case <-ctx.Done():
			log.Println("Config polling stopped")
			return
		}
	}
}

func (m *Manager) fetchAndUpdate(onUpdate func([]models.BackupJob)) {
	config, err := m.client.GetAgentConfig()
	if err != nil {
		log.Printf("Failed to fetch config: %v", err)
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if config.ConfigVersion != m.configHash {
		log.Printf("Config changed: %s -> %s", m.configHash, config.ConfigVersion)
		m.configHash = config.ConfigVersion
		m.currentConfig = config

		log.Printf("Loaded %d backup jobs", len(config.BackupJobs))

		if onUpdate != nil {
			onUpdate(config.BackupJobs)
		}
	} else {
		log.Printf("Config unchanged (version: %s)", m.configHash)
	}
}

func (m *Manager) GetCurrentConfig() *models.AgentConfigResponse {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.currentConfig
}
