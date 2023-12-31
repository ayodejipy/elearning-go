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

// POST /signup
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
		return
	}

	// set default role to user and update password to hashed password
	user.Role = models.UserRole 
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

// POST /login
func (con *Handler) Login(w http.ResponseWriter, r *http.Request) {
	body := &models.User{} // var to have request body parsed json

	// read body from request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to read request body.")
		return
	}
	// get user from database
	var user models.User
	if err := con.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
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
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Failed to create token")
		return
	}

	// we have tokenString here: save tokenstring to cookie
	cookie := http.Cookie{
		Name: "Authorization",
		Value: tokenString,
		SameSite: http.SameSiteLaxMode,
		MaxAge: 3600 * 24 * 30,
		Secure: false, // true in prod
		Path: "/",
		Domain: "",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)


	// send jwt token back
	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "User logged in successfully.",
	})

}

// GET /validate
func (con *Handler) Validate(w http.ResponseWriter, r *http.Request) {
	// read our user on the context
	user := r.Context().Value("user").(models.User)
	
	// send jwt token back
	helpers.RespondWithJSON(w, http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "Request verified successfully.",
		Data: user.Username,
	})
}

// GET: /becomeatutor
func (con *Handler) BecomeTutor(w http.ResponseWriter, r *http.Request) {
	// make an empty tutor struct
	tutor := models.Tutors{}
	// get signed in user object from req context
	user := r.Context().Value("user").(models.User)
	
	// add data to tutor's table
	if err := con.DB.Model(&user).Updates(models.User{
		Role: models.TutorRole,
	}).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Failed to update user role")
		return
	}

	// update tutor details
	tutor.UserId = user.ID
	tutor.CreatedAt = time.Now()

	// add data to tutor's table
	if err := con.DB.Create(&tutor).Error; err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Failed to create tutor")
		return
	}


	helpers.RespondWithJSON(w, http.StatusCreated, helpers.BaseResponse{
		Success: true,
		Message: "Tutor created successfully",
	})
}