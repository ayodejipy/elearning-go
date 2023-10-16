package services

import (
	"fmt"
	"time"

	"github.com/ayodejipy/elearning-go/internal/database"
	"gorm.io/gorm"
)


type Course struct {
	gorm.Model

	Title string `json:"title"`
	Description string `json:"description"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Price float64 `json:"price"`
	TutorID int `json:"tutor_id"`
}

type Handler struct {
	*database.DBConn
}

func CourseService(db *database.DBConn) *Handler {
	return &Handler{db}
}


func (con *Handler) GetAllCourses() (*[]Course, error) {
	var courses *[]Course

	if err := con.DB.Find(&courses).Error; err != nil {
		fmt.Errorf("error fetching all courses: %v", err)
		return nil, err
	}
	
	// return data
	return courses, nil
}

// // POST: Adds a new course to database
// func (con *Handler) AddCourse(w http.ResponseWriter, r *http.Request) {
// 	// make an empty course struct/object
// 	course := models.Course{}

// 	err := json.NewDecoder(r.Body).Decode(&course)
// 	if err != nil {
// 		helpers.RespondWithError(w, http.StatusInternalServerError, "Invalid request body")
// 		fmt.Printf("Cannot decode request: %v \n", err)
// 		return 
// 	}

// 	if notZero := time.Time.IsZero(course.CreatedAt); notZero {
// 		// course.StartDate = time.Now()
// 		// course.EndDate = time.Now()
// 		course.CreatedAt = time.Now()
// 		course.UpdatedAt = time.Now()
// 	}

// 	if err := con.DB.Create(&course).Error; err != nil {
// 		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to create new course")
// 		return
// 	}

// 	helpers.RespondWithJSON(w, http.StatusCreated, helpers.BaseResponse{
// 		Success: true,
// 		Message: "Course added successfully.",
// 	})
// }

// // GET: Get course by id
// func (con *Handler) GetCourseById(w http.ResponseWriter, r *http.Request) {
// 	course := models.Course{}

// 	// get course id from url param
// 	courseId := chi.URLParam(r, "id")

// 	// use id to query db
// 	if err := con.DB.First(&course, "ID = ?", courseId).Error; err != nil {
// 		helpers.RespondWithError(w, http.StatusBadRequest, "Cannot find course with ID")
// 		return
// 	}

// 	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
// 		Success: true,
// 		Message: "Course fetched successfully",
// 		Data: course,
// 	})

// }

// // UPDATE: Update course by id
// func (con *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	
// }