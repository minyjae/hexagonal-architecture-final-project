package services

import (
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
)

type listQueueService struct {
	repo        repoPort.ListQueueRepository
	mappingRepo repoPort.StatusMappingRepository
}

func NewListQueueServiceImpl(r repoPort.ListQueueRepository, m repoPort.StatusMappingRepository) *listQueueService {
	return &listQueueService{repo: r, mappingRepo: m}
}

func (r *listQueueService) Create(q *entities.ListQueue) (*entities.ListQueue, error) {
	if err := r.repo.Create(q); err != nil {
		return nil, fmt.Errorf("create queue: %w", err)
	}

	return q, nil
}

func (s *listQueueService) Update(q *entities.ListQueue) (*entities.ListQueue, error) {
	if err := s.repo.Update(q); err != nil {
		return nil, fmt.Errorf("update queue: %w", err)
	}
	return q, nil
}

func (s *listQueueService) UpdateStaffStatus(queueID, staffStatusID uint) (*entities.ListQueue, error) {
	q, err := s.repo.FindByID(queueID)
	if err != nil {
		return nil, err
	}

	// หา mapping → userStatus
	m, err := s.mappingRepo.FindByStaffStatusID(staffStatusID)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, fmt.Errorf("no mapping for staffStatus=%d", staffStatusID)
	}

	q.StaffStatusID = staffStatusID
	q.UserStatusID = m.UserStatusID

	if err := s.repo.Update(q); err != nil {
		return nil, err
	}
	return q, nil
}

func (r *listQueueService) FindByID(id uint) (*entities.ListQueue, error) {
	return r.repo.FindByID(id)
}

func (r *listQueueService) FindAll() ([]entities.ListQueue, error) {
	return r.repo.FindAll()
}

func (r *listQueueService) FindByStaffStatus(staffStatusID uint) ([]entities.ListQueue, error) {
	return r.repo.FindByStaffStatus(staffStatusID)
}

func (r *listQueueService) FindByUserStatus(userStatusID uint) ([]entities.ListQueue, error) {
	return r.repo.FindByUserStatus(userStatusID)
}
