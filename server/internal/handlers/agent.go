package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wimwenigerkind/backup-manager/server/internal/database"
	"github.com/wimwenigerkind/backup-manager/server/internal/dto"
	"github.com/wimwenigerkind/backup-manager/server/internal/models"
)

type AgentHandler struct{}

func NewAgentHandler() *AgentHandler {
	return &AgentHandler{}
}

func (h *AgentHandler) CreateAgent(c *gin.Context) {
	var req dto.CreateAgentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	agent := models.Agent{
		Name: req.Name,
		IP:   req.IP,
	}

	if err := database.DB.Create(&agent).Error; err != nil {
		log.Printf("Error creating agent: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create agent", "details": err.Error()})
		return
	}

	response := dto.AgentResponse{
		ID:   agent.ID.String(),
		Name: agent.Name,
		IP:   agent.IP,
	}

	c.JSON(http.StatusCreated, response)
}
