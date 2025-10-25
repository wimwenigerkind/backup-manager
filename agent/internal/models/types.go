package models

type AgentConfigResponse struct {
	ConfigVersion string      `json:"config_version"`
	Agent         AgentInfo   `json:"agent"`
	BackupJobs    []BackupJob `json:"backup_jobs"`
}

type AgentInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}

type BackupJob struct {
	ID            string         `json:"id"`
	Interval      int            `json:"interval"`
	Source        string         `json:"source"`
	BackupTargets []BackupTarget `json:"backup_targets"`
}

type BackupTarget struct {
	ID         string `json:"id"`
	TargetType string `json:"target_type"`
	Path       string `json:"path"`
}
