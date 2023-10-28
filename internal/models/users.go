package models

const (
	UserRole = "user"
	TutorRole = "tutor"
)

type User struct {
	UserId
	BaseModel

	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"-"`
	Role string `json:"role"`
}