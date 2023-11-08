package models

import (
	"time"

	"github.com/google/uuid"
)

type Enrollment struct {
	DataID
	
	UserId *uuid.UUID
	CourseId int64
	EnrollmentDate time.Time
}