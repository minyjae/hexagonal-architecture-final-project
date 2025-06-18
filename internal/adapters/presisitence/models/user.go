package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

// User table
type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Email       string    `gorm:"unique;not null" json:"email"`
	Faculty     string    `gorm:"not null" json:"faculty"`
	FirstName   string    `gorm:"not null" json:"first_name"`
	LastName    string    `gorm:"not null" json:"last_name"`
	PhoneNumber string    `gorm:"not null" json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *User) ToEntity() *entities.User {
	return &entities.User{
		ID:          u.ID,
		Email:       u.Email,
		Faculty:     u.Faculty,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		PhoneNumber: u.PhoneNumber,
	}
}

func (u *User) FromEntity(entity *entities.User) {
	u.ID = entity.ID
	u.Email = entity.Email
	u.Faculty = entity.Faculty
	u.FirstName = entity.FirstName
	u.LastName = entity.LastName
	u.PhoneNumber = entity.PhoneNumber
}
