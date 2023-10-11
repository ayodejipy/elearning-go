package handlers

import (
	"fmt"
	"net/http"

	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/helpers"
	"github.com/ayodejipy/elearning-go/internal/models"
)

type Handler struct {
	db *database.DBConn
}

func CoursesHandler(db *database.DBConn) *Handler {
	return &Handler{db}
}

// Fetch all courses
func (con *Handler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses := []models.Courses{}

	if err := con.db.Find(&courses).Error; err != nil {
		fmt.Errorf("error fetching all courses: %v", err)
		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to fetch all courses")
	}
	
	// respond with the data
	helpers.RespondWithJSON(w, http.StatusOK, courses)
}

// Adds a new course to database
func (con *Handler) AddCourse(w http.ResponseWriter, r *http.Request) {
	return
}

func (con *Handler) GetCourseById(w http.ResponseWriter, r *http.Request) {
	return
}