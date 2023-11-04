package handlers

import "github.com/ayodejipy/elearning-go/internal/database"


func tutorHandler(db *database.DBConn) *Handler {
	return &Handler{db}
}

// UPDATE /course/