package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayodejipy/elearning-go/internal/database"
	"github.com/ayodejipy/elearning-go/internal/helpers"
	"github.com/ayodejipy/elearning-go/internal/models"
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

	fmt.Printf("Decoded user: %v", user)

	// hash user password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Cannot hash password.")
	}

	// update password to hashed password
	user.Password = string(hash)

	// save to database
	if err := con.DB.Create(&user).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Failed to create user.")
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, helpers.BaseResponse{
		Success: true,
		Message: "User created successfully",
	})
}

func (con *Handler) Login(w http.ResponseWriter, r *http.Request) {}
