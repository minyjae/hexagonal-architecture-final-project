package repositories

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUser() ([]*entities.User, error)
	RemoveUser(id uint) (*entities.User, error)
	GetUserByFaculty(faculty string) (*entities.User, error)
	UpdateUserFaculty(id uint, faculty string) (*entities.User, error)
}
