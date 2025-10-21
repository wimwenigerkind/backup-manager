package models

import "github.com/google/uuid"

type Agent struct {
	ID   uuid.UUID `gorm:"type:char(36);primary_key;" json:"id"`
	Name string    `gorm:"type:varchar(100);not null;" json:"name"`
	IP   string    `gorm:"type:varchar(45);not null;" json:"ip"`
}
