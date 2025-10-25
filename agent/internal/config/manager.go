package config

import (
	"sync"

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

func (m *Manager) GetCurrentConfig() *models.AgentConfigResponse {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.currentConfig
}
