package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

// UserStatus table
type UserStatus struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Status    string      `gorm:"not null;unique" json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Lists     []ListQueue `gorm:"foreignKey:UsersStatusID"`
}

func (u *UserStatus) ToEntity() *entities.UserStatus {
	return &entities.UserStatus{
		ID:     u.ID,
		Status: u.Status,
	}
}

func (u *UserStatus) FromEntity(entity *entities.UserStatus) {
	u.ID = entity.ID
	u.Status = entity.Status
}
