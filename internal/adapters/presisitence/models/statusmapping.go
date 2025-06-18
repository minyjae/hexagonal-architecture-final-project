package models

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

// StatusMapping table
type StatusMapping struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	StaffStatusID uint        `gorm:"not null;index" json:"staff_status_id"`
	StaffStatus   StaffStatus `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"staff_status"`
	UserStatusID  uint        `gorm:"not null;index" json:"user_status_id"`
	UserStatus    UserStatus  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_status"`
}

func (s *StatusMapping) ToEntity() *entities.StatusMapping {
	return &entities.StatusMapping{
		ID:            s.ID,
		StaffStatusID: s.StaffStatusID,
		UserStatusID:  s.UserStatusID,
	}
}

func (s *StatusMapping) FromEntity(entity *entities.StatusMapping) {
	s.ID = entity.ID
	s.StaffStatusID = entity.StaffStatusID
	s.UserStatusID = entity.UserStatusID
}
