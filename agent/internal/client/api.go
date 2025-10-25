package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/wimwenigerkind/backup-manager/agent/internal/models"
)

type APIClient struct {
	ServerUrl string
	AgentID   string
	client    *http.Client
}

func NewAPIClient(serverURL string, agentID string) *APIClient {
	return &APIClient{
		ServerUrl: serverURL,
		AgentID:   agentID,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *APIClient) GetAgentConfig() (*models.AgentConfigResponse, error) {
	url := fmt.Sprintf("%s/api/v1/agents/%s/config", c.ServerUrl, c.AgentID)

	response, err := c.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get agent config: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var config models.AgentConfigResponse
	if err := json.NewDecoder(response.Body).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}
	return &config, nil
}
