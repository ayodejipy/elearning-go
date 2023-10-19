package routes

import (
	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, db *database.DBConn) {
	// bind handler to database
	courseHandler := handlers.CoursesHandler(db)
	userHandler := handlers.UserHandler(db)
	

	// declare routes
	r.Get("/", handlers.GetHelloHandler)

	r.Get("/courses", courseHandler.GetAllCourses)
	r.Post("/course/create", courseHandler.AddCourse)
	r.Get("/course/{id}", courseHandler.GetCourseById)
	r.Put("/course/{course_id}", courseHandler.UpdateCourse)
	r.Delete("/course/{course_id}", courseHandler.DeleteCourse)
	
	// Auth routes
	r.Post("/auth/signup", userHandler.RegisterUser)
}
