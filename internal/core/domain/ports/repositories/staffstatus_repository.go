package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type StaffStatusRepository interface {
	Create(status *entities.StaffStatus) error
	Delete(id uint) error
	FindAll() ([]entities.StaffStatus, error)

	// Mapping‑related
	FindMapping(staffStatusID uint) (*entities.StatusMapping, error)
	UpsertMapping(mapping *entities.StatusMapping) error
	FindUserStatusByStaff(staffStatusID uint) (*entities.UserStatus, error)
}
