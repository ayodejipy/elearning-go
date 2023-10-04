package routes

import (
	"net/http"

	"github.com/ayodejipy/elearning-go/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// declare routes
	r.Get("/", handlers.GetHelloHandler)


	return r // return the router instance
}