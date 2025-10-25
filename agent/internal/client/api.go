package client

import "net/http"

type APIClient struct {
	ServerUrl string
	AgentID   string
	client    *http.Client
}
