package domain

import (
	"go-repaso/internal/models"
	"time"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusQueued     TaskStatus = "queued"
	StatusProcessing TaskStatus = "processing"
	StatusDone       TaskStatus = "done"
	StatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	UserID       uint        `gorm:"not null;index" json:"user_id"`
	User         models.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Title        string      `gorm:"type:varchar(255);not null" json:"title"`
	Description  string      `gorm:"type:text" json:"description"`
	Status       TaskStatus  `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	Result       *string     `gorm:"type:text" json:"result,omitempty"`
	ErrorMessage *string     `gorm:"type:text" json:"error_message,omitempty"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
