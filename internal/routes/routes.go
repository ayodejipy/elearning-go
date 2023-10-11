package routes

import (
	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, db *database.DBConn) {
	// bind handler to database
	courseHandler := handlers.CoursesHandler(db)

	// declare routes
	r.Get("/", handlers.GetHelloHandler)

	r.Get("/courses", courseHandler.GetAllCourses)
}
