package models

import (
	"time"
)

type Course struct {
	DataID
	BaseModel

	Title string `json:"title"`
	Description string `json:"description"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Price float64 `json:"price"`
	TutorID int `json:"tutor_id"`
}