package models

import "github.com/google/uuid"

type Tutors struct {
	DataID
	
	UserId *uuid.UUID
	BaseModel

}