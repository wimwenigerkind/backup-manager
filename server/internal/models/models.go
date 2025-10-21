package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}

type Agent struct {
	BaseModel
	Name       string      `gorm:"type:varchar(100);not null" json:"name"`
	IP         string      `gorm:"type:varchar(45);not null" json:"ip"`
	BackupJobs []BackupJob `gorm:"foreignKey:AgentID" json:"backup_jobs,omitempty"`
}

type BackupJob struct {
	BaseModel
	AgentID       uuid.UUID      `gorm:"type:char(36);not null;index" json:"agent_id"`
	Agent         *Agent         `gorm:"foreignKey:AgentID" json:"agent,omitempty"`
	Interval      int            `gorm:"not null" json:"interval"`
	Source        string         `gorm:"type:varchar(255);not null" json:"source"`
	BackupTargets []BackupTarget `gorm:"foreignKey:BackupJobID" json:"backup_targets,omitempty"`
}

type BackupTarget struct {
	BaseModel
	BackupJobID uuid.UUID  `gorm:"type:char(36);not null;index" json:"backup_job_id"`
	BackupJob   *BackupJob `gorm:"foreignKey:BackupJobID" json:"backup_job,omitempty"`
	TargetType  string     `gorm:"type:varchar(50);not null" json:"target_type"`
	Path        string     `gorm:"type:varchar(255);not null" json:"path"`
}
