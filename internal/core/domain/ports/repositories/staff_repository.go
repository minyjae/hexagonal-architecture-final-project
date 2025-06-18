package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type StaffRepository interface {
	CreateStaff(req *entities.Staff) (*entities.Staff, error)
	RemoveStaff(id uint) (*entities.Staff, error)
	GetStaffByEmail(email string) (*entities.Staff, error)
}
