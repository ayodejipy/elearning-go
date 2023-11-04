package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/helpers"
	"github.com/ayodejipy/elearning-go/internal/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CoursesHandler(db *database.DBConn) *Handler {
	return &Handler{db}
}

// GET: /courses: Fetch all courses
func (con *Handler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses := []models.Course{}

	if err := con.DB.Find(&courses).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to fetch all courses")
		return
	}
	
	// respond with the data
	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "Courses fetched successfully.",
		Data: courses,
	})
}

// POST: /course/create Adds a new course to database
func (con *Handler) AddCourse(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)

	// make an empty course struct/object
	course := models.Course{}

	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Invalid request body")
		fmt.Printf("Cannot decode request: %v \n", err)
		return 
	}
	// updated course data
	if notZero := time.Time.IsZero(course.CreatedAt); notZero {
		course.CreatedAt = time.Now()
		course.UpdatedAt = time.Now()
	}
	course.TutorID = *user.ID

	if err := con.DB.Create(&course).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to create new course")
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, helpers.BaseResponse{
		Success: true,
		Message: "Course added successfully.",
	})
}

// GET: /course/{id}
func (con *Handler) GetCourseById(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	// get course id from url param
	courseId := chi.URLParam(r, "id")

	// use id to query db
	if err := con.DB.First(&course, "ID = ?", courseId).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Cannot find course with ID")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "Course fetched successfully",
		Data: course,
	})

}

// UPDATE: PUT /course/{couse_id}
func (con *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}
	// get id from url param
	id := chi.URLParam(r, "course_id")

	// decode request body into json
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Cannot parse request as JSON")
		return 
	}
	
	user := r.Context().Value("user").(models.User)
	
	// UPDATE course in database; ensure it's the user's course
	if err := con.DB.Model(&course).Clauses(clause.Returning{}).Where("id = ? AND tutor_id = ?", id, user.ID).Updates(models.Course{
		Title: course.Title,
		Description: course.Description,
		Price: course.Price,
		TutorID: course.TutorID,
		StartDate: course.StartDate,
		EndDate: course.EndDate,
	}).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to update courses")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "Course updated successfully",
		Data: course,
	})
}

// DELETE /course/{course_id}
func (con *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}
	// get id from url param
	id := chi.URLParam(r, "course_id")

	if err := con.DB.Delete(&course, id).Error; err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Unable to delete delete course")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "Course deleted successfully",
	})
}

// GET /courses/{tutor_id}
func (con *Handler) GetCoursesByTutor(w http.ResponseWriter, r *http.Request) {
	// get course id from the request
	tutor_id := chi.URLParam(r, "tutor_id")
	
	// find tutor
	tutor := models.Tutors{}
	if err := con.DB.Where("user_id = ?", tutor_id).Find(&tutor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.RespondWithError(w, http.StatusBadRequest, "Tutor does not exist.")
			return
		}
		// for other type of errors
		helpers.RespondWithError(w, http.StatusBadRequest, "Failed to get tutor.")
		return
	}

	// query database and return all the courses belonging to tutor
	courses := []models.Course{}
	if err := con.DB.Where("tutor_id = ?", tutor_id).Find(&courses).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Courses not found.")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "Courses fetched successfully",
		Data: courses,
	})
}