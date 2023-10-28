package routes

import (
	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/handlers"
	"github.com/ayodejipy/elearning-go/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRoutes(r *chi.Mux, db *database.DBConn) {
	// setup cors options
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// bind handler to database
	courseHandler := handlers.CoursesHandler(db)
	userHandler := handlers.UserHandler(db)
	

	// declare routes
	r.Get("/", handlers.GetHelloHandler)
	
	// Auth routes
	r.Post("/auth/signup", userHandler.RegisterUser)
	r.Post("/auth/signin", userHandler.Login)
	

	// Private Routes
    // Require Authentication
    r.Group(func(r chi.Router) {
		r.Use(middlewares.VerifyAuth(db))
		r.Use(middlewares.TutorMiddleware)

		r.Get("/validate", userHandler.Validate)

		r.Get("/courses", courseHandler.GetAllCourses)
		r.Post("/course/create", courseHandler.AddCourse)
		r.Get("/course/{id}", courseHandler.GetCourseById)
		r.Put("/course/{course_id}", courseHandler.UpdateCourse)
		r.Delete("/course/{course_id}", courseHandler.DeleteCourse)

		r.Get("/user/becomeTutor", userHandler.BecomeTutor)
    })
}
