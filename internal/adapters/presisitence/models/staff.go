package models

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type Staff struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Email string `gorm:"unique;not null" json:"email"`
}

func (s *Staff) ToEntity() *entities.Staff {
	return &entities.Staff{
		ID:    s.ID,
		Email: s.Email,
	}
}

func (s *Staff) FromEntity(entity *entities.Staff) {
	s.ID = entity.ID
	s.Email = entity.Email
}
