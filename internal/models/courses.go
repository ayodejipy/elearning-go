package models

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	DataID
	BaseModel

	Title string `json:"title"`
	Description string `json:"description"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Price float64 `json:"price"`
	TutorID uuid.UUID `json:"tutor_id"`
}