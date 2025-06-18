package services

import (
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
)

type staffStatusService struct {
	repo repoPort.StaffStatusRepository
}

func NewStaffStatusServiceImpl(r repoPort.StaffStatusRepository) *staffStatusService {
	return &staffStatusService{repo: r}
}

func (s *staffStatusService) Create(st *entities.StaffStatus) error {
	return s.repo.Create(st)
}

func (s *staffStatusService) Remove(id uint) error {
	return s.repo.Delete(id)
}

func (s *staffStatusService) List() ([]entities.StaffStatus, error) {
	return s.repo.FindAll()
}

func (s *staffStatusService) BindMappings(mappings []entities.StatusMapping) error {
	for _, m := range mappings {
		exist, err := s.repo.FindMapping(m.StaffStatusID)
		if err != nil {
			return err
		}

		if exist == nil {
			if err := s.repo.UpsertMapping(&m); err != nil {
				return fmt.Errorf("create mapping: %w", err)
			}
			continue
		}

		// update เฉพาะกรณี ID เปลี่ยน
		if exist.UserStatusID != m.UserStatusID {
			exist.UserStatusID = m.UserStatusID
			if err := s.repo.UpsertMapping(exist); err != nil {
				return fmt.Errorf("update mapping: %w", err)
			}
		}
	}
	return nil
}

func (s *staffStatusService) GetUserStatus(staffStatusID uint) (*entities.UserStatus, error) {
	return s.repo.FindUserStatusByStaff(staffStatusID)
}
