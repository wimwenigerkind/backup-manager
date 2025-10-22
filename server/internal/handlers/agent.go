package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"strings"

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

func (h *AgentHandler) GetAgents(c *gin.Context) {
	var agents []models.Agent
	if err := database.DB.
		Preload("BackupJobs").
		Preload("BackupJobs.BackupTargets").
		Find(&agents).Error; err != nil {
		log.Printf("Error fetching agents: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch agents", "details": err.Error()})
		return
	}

	response := make([]dto.AgentResponse, 0, len(agents))
	for _, agent := range agents {
		backupJobs := make([]dto.BackupJobResponse, 0, len(agent.BackupJobs))
		for _, job := range agent.BackupJobs {
			backupTargets := make([]dto.BackupTargetResponse, 0, len(job.BackupTargets))
			for _, target := range job.BackupTargets {
				backupTargets = append(backupTargets, dto.BackupTargetResponse{
					ID:         target.ID.String(),
					TargetType: target.TargetType,
					Path:       target.Path,
				})
			}

			backupJobs = append(backupJobs, dto.BackupJobResponse{
				ID:            job.ID.String(),
				Interval:      job.Interval,
				Source:        job.Source,
				BackupTargets: backupTargets,
			})
		}

		response = append(response, dto.AgentResponse{
			ID:         agent.ID.String(),
			Name:       agent.Name,
			IP:         agent.IP,
			BackupJobs: backupJobs,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *AgentHandler) CreateBackupJob(c *gin.Context) {
	agentID := c.Param("id")

	var agent models.Agent
	if err := database.DB.First(&agent, "id = ?", agentID).Error; err != nil {
		log.Printf("Agent not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}

	var req dto.CreateBackupJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	backupJob := models.BackupJob{
		AgentID:  agent.ID,
		Interval: req.Interval,
		Source:   req.Source,
	}

	if err := database.DB.Create(&backupJob).Error; err != nil {
		log.Printf("Error creating backup job: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create backup job", "details": err.Error()})
		return
	}

	response := dto.BackupJobResponse{
		ID:       backupJob.ID.String(),
		Interval: backupJob.Interval,
		Source:   backupJob.Source,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AgentHandler) DeleteBackupJob(c *gin.Context) {
	jobID := c.Param("id")

	var job models.BackupJob
	if err := database.DB.First(&job, "id = ?", jobID).Error; err != nil {
		log.Printf("Backup job not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Backup job not found"})
		return
	}

	if err := database.DB.Delete(&job).Error; err != nil {
		log.Printf("Error deleting backup job: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete backup job", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AgentHandler) CreateBackupTarget(c *gin.Context) {
	jobID := c.Param("id")

	var job models.BackupJob
	if err := database.DB.First(&job, "id = ?", jobID).Error; err != nil {
		log.Printf("Backup job not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Backup job not found"})
		return
	}

	var req dto.CreateBackupTargetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	target := models.BackupTarget{
		BackupJobID: job.ID,
		TargetType:  req.TargetType,
		Path:        req.Path,
	}

	if err := database.DB.Create(&target).Error; err != nil {
		log.Printf("Error creating backup target: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create backup target", "details": err.Error()})
		return
	}

	response := dto.BackupTargetResponse{
		ID:         target.ID.String(),
		TargetType: target.TargetType,
		Path:       target.Path,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AgentHandler) DeleteBackupTarget(c *gin.Context) {
	targetID := c.Param("id")

	var target models.BackupTarget
	if err := database.DB.First(&target, "id = ?", targetID).Error; err != nil {
		log.Printf("Backup target not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Backup target not found"})
		return
	}

	if err := database.DB.Delete(&target).Error; err != nil {
		log.Printf("Error deleting backup target: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete backup target", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AgentHandler) GetAgentConfig(c *gin.Context) {
	agentID := c.Param("id")

	var agent models.Agent
	if err := database.DB.
		Preload("BackupJobs").
		Preload("BackupJobs.BackupTargets").
		First(&agent, "id = ?", agentID).Error; err != nil {
		log.Printf("Agent not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}

	// Build backup jobs response
	backupJobs := make([]dto.BackupJobResponse, 0, len(agent.BackupJobs))
	for _, job := range agent.BackupJobs {
		backupTargets := make([]dto.BackupTargetResponse, 0, len(job.BackupTargets))
		for _, target := range job.BackupTargets {
			backupTargets = append(backupTargets, dto.BackupTargetResponse{
				ID:         target.ID.String(),
				TargetType: target.TargetType,
				Path:       target.Path,
			})
		}

		backupJobs = append(backupJobs, dto.BackupJobResponse{
			ID:            job.ID.String(),
			Interval:      job.Interval,
			Source:        job.Source,
			BackupTargets: backupTargets,
		})
	}

	// Generate config version hash based on agent data
	configVersion := generateConfigHash(agent)

	response := dto.AgentConfigResponse{
		ConfigVersion: configVersion,
		Agent: dto.AgentInfo{
			ID:   agent.ID.String(),
			Name: agent.Name,
			IP:   agent.IP,
		},
		BackupJobs: backupJobs,
	}

	c.JSON(http.StatusOK, response)
}

func generateConfigHash(agent models.Agent) string {
	// Create a hash based on agent data and all jobs/targets timestamps
	var builder strings.Builder

	// Include agent data
	builder.WriteString(agent.ID.String())
	builder.WriteString(agent.UpdatedAt.Format("20060102150405"))

	// Include all backup jobs and their targets
	for _, job := range agent.BackupJobs {
		builder.WriteString(job.ID.String())
		builder.WriteString(job.UpdatedAt.Format("20060102150405"))

		for _, target := range job.BackupTargets {
			builder.WriteString(target.ID.String())
			builder.WriteString(target.UpdatedAt.Format("20060102150405"))
		}
	}

	// Generate SHA-256 hash
	hash := sha256.Sum256([]byte(builder.String()))
	return hex.EncodeToString(hash[:])
}
