package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type ListQueueService interface {
	Create(q *entities.ListQueue) (*entities.ListQueue, error)
	Update(q *entities.ListQueue) (*entities.ListQueue, error)

	UpdateStaffStatus(queueID, staffStatusID uint) (*entities.ListQueue, error)
	UpdatePriority(queueID uint, priority int) (*entities.ListQueue, error)

	GetAll() ([]entities.ListQueue, error)
	GetByID(id uint) (*entities.ListQueue, error)
	GetByStaffStatus(staffStatusID uint) ([]entities.ListQueue, error)
	GetByUserStatus(userStatusID uint) ([]entities.ListQueue, error)
}
