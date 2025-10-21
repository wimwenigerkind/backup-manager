package dto

type CreateAgentRequest struct {
	Name string `json:"name" binding:"required"`
	IP   string `json:"ip" binding:"required"`
}

type AgentResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}
