package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/helpers"
	"github.com/ayodejipy/elearning-go/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyAuth(db *database.DBConn) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get cookie off the request r context
			tokenString, err := r.Cookie("Authorization")
			if err != nil {
				fmt.Println(err)
				helpers.RespondWithError(w, http.StatusUnauthorized, "User not authorized: getting cookie")
				return
			}

			// decode and validate token
			token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// secret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// check expiration date
				if float64(time.Now().Unix()) > claims["exp"].(float64) {
					helpers.RespondWithError(w, http.StatusUnauthorized, "JWT expired.")
					return
				}

				// Fetch the user with ID
				var user models.User
				
				if err := db.DB.Where("id = ?", claims["sub"]).First(&user).Error; err != nil {
					helpers.RespondWithError(w, http.StatusBadRequest, "Failed to fetch user")
					return
				}

				// create new context from `r` request context, and update our request
				ctx := context.WithValue(r.Context(), "user", user)

				// continue request
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				helpers.RespondWithError(w, http.StatusUnauthorized, "User not authorized: token claim failed")
				return
			}

		})
	}
}

