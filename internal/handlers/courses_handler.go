package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/helpers"
	"github.com/ayodejipy/elearning-go/internal/models"
)

type Handler struct {
	*database.DBConn
}

func CoursesHandler(db *database.DBConn) *Handler {
	return &Handler{db}
}

// GET: Fetch all courses
func (con *Handler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses := []models.Course{}

	if err := con.DB.Find(&courses).Error; err != nil {
		fmt.Errorf("error fetching all courses: %v", err)
		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to fetch all courses")
	}
	
	// respond with the data
	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "Courses fetched successfully.",
		Data: courses,
	})
}

// POST: Adds a new course to database
func (con *Handler) AddCourse(w http.ResponseWriter, r *http.Request) {
	// make an empty course struct/object
	course := models.Course{}

	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Invalid request body")
		fmt.Printf("Cannot decode request: %v \n", err)
		return 
	}

	if notZero := time.Time.IsZero(course.CreatedAt); notZero {
		// course.StartDate = time.Now()
		// course.EndDate = time.Now()
		course.CreatedAt = time.Now()
		course.UpdatedAt = time.Now()
	}

	if err := con.DB.Create(&course).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to create new course")
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, helpers.BaseResponse{
		Success: true,
		Message: "Course added successfully.",
	})
}

// GET: Get course by id
func (con *Handler) GetCourseById(w http.ResponseWriter, r *http.Request) {
	return
}