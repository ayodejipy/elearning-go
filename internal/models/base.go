package models

import (
	"time"

	"github.com/google/uuid"
)

// gorm.Model definition
type BaseModel struct {
  ID        uuid.UUID           `gorm:"primary_key" json:"id,omitempty"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
