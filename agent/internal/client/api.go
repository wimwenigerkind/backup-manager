package client

import (
	"net/http"
	"time"
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
