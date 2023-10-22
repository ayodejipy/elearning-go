package models

import (
	"time"

	"github.com/google/uuid"
)

type DataID struct {
  ID  int64           `gorm:"primary_key" json:"id,omitempty"`
}
type UserId struct {
  ID        *uuid.UUID           `gorm:"primary_key" json:"id,omitempty"`
}

// gorm.Model definition
type BaseModel struct {
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
