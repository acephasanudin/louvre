package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Example represents an example entity
type Example struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Description string       `gorm:"type:text" json:"description"`
	Status    string         `gorm:"default:'active'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName returns the table name for the Example model
func (Example) TableName() string {
	return "examples"
}