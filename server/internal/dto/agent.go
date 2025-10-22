package dto

type CreateAgentRequest struct {
	Name string `json:"name" binding:"required"`
	IP   string `json:"ip" binding:"required"`
}

type BackupTargetResponse struct {
	ID         string `json:"id"`
	TargetType string `json:"target_type"`
	Path       string `json:"path"`
}

type BackupJobResponse struct {
	ID            string                 `json:"id"`
	Interval      int                    `json:"interval"`
	Source        string                 `json:"source"`
	BackupTargets []BackupTargetResponse `json:"backup_targets,omitempty"`
}

type AgentResponse struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	IP         string              `json:"ip"`
	BackupJobs []BackupJobResponse `json:"backup_jobs,omitempty"`
}

type CreateBackupJobRequest struct {
	Interval int    `json:"interval" binding:"required,min=1"`
	Source   string `json:"source" binding:"required"`
}
