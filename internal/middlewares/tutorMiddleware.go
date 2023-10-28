package middlewares

import (
	"net/http"

	"github.com/ayodejipy/elearning-go/internal/helpers"
	"github.com/ayodejipy/elearning-go/internal/models"
)


func TutorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get user details from the req context
		user := r.Context().Value("user").(models.User)

		if user.Role == models.TutorRole {
			//proceed with request
			next.ServeHTTP(w, r);
		} else {
			helpers.RespondWithError(w, http.StatusUnauthorized, "User cannot perform this action")
		}
	})
}