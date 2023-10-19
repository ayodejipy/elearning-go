package models

import (
	"time"
)

// gorm.Model definition
type BaseModel struct {
  ID        uint           `gorm:"primary_key" json:"id,omitempty"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
