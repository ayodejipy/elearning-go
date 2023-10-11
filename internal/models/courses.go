package models

import (
	"time"

	"gorm.io/gorm"
)

type Courses struct {
	gorm.Model

	Title string `json:"title"`
	Description string `json:"description"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Price float64 `json:"price"`
	TutorID int `json:"tutor_id"`
}