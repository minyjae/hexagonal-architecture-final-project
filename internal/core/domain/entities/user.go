package entities

import (
	"time"
)

type User struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	Faculty     string    `json:"faculty"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserLoginRequest struct {
	Email string `json:"email"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
}
