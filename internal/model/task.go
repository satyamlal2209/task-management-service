package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Status      string         `gorm:"type:varchar(20);not null" json:"status"` // Pending, InProgress, Completed
	DueDate     time.Time      `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
