package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type StaffStatus struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Status    string      `gorm:"not null;unique" json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Lists     []ListQueue `gorm:"foreignKey:StaffStatusID"`
}

func (s *StaffStatus) ToEntity() *entities.StaffStatus {
	return &entities.StaffStatus{
		ID:     s.ID,
		Status: s.Status,
	}
}

func (s *StaffStatus) FromEntity(entity *entities.StaffStatus) {
	s.ID = entity.ID
	s.Status = entity.Status
}
