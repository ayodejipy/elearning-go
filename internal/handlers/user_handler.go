package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/helpers"
	"github.com/ayodejipy/elearning-go/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func UserHandler(db *database.DBConn) *Handler {
	return &Handler{db}
}

func (con *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	// parse body into json
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to read request body.")
		return
	}

	// hash user password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Cannot hash password.")
	}

	// update password to hashed password
	user.Password = string(hash)

	// save to database
	if err := con.DB.Omit("id").Create(&user).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Failed to create user.")
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, helpers.BaseResponse{
		Success: true,
		Message: "User created successfully",
	})
}

func (con *Handler) Login(w http.ResponseWriter, r *http.Request) {
	body := &models.User{} // var to have request body parsed json

	// read body from request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to read request body.")
		return
	}
	// get user from database
	var user *models.User
	if err := con.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "User with email not found.")
		return
	}

	// compare user password and stored password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid email or password")
		return
	}

	// Generate a new jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Failed to create token")
		return
	}


	// send jwt token back
	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Message: "Login successful",
		Data: map[string]string{
			"token": tokenString,
		},
	})

}
