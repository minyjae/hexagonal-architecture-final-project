package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type UserStatusRepository interface {
	CreateUserStatus(req *entities.UserStatus) (*entities.UserStatus, error)
	RemoveUserStatus(id uint) error
	GetUserStatus() (*[]entities.UserStatus, error)
}
