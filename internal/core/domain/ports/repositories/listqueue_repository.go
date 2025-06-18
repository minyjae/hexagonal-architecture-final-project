package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type ListQueueRepository interface {
	Create(q *entities.ListQueue) error
	Update(q *entities.ListQueue) error
	FindByID(id uint) (*entities.ListQueue, error)

	FindAll() ([]entities.ListQueue, error)
	FindByStaffStatus(staffStatusID uint) ([]entities.ListQueue, error)
	FindByUserStatus(userStatusID uint) ([]entities.ListQueue, error)
}
