package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type StaffStatusService interface {
	Create(status *entities.StaffStatus) (*entities.StaffStatus, error)
	Remove(id uint) error
	GetUserStatus(staffStatID uint) (*entities.UserStatus, error)
	GetStaffStatus() (*[]entities.StaffStatus, error)
	BindStatusMappings(mappings *[]entities.StatusMapping) error
	List() ([]entities.StaffStatus, error)
}
