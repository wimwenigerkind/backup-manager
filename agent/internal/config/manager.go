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
