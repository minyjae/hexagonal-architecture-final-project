package repositories

import (
	"errors"
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presisitence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type StaffStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewStaffStatusRepository(db *gorm.DB) *StaffStatusRepositoryImpl {
	return &StaffStatusRepositoryImpl{db: db}
}

func (r *StaffStatusRepositoryImpl) CreateStaffStatus(req *entities.StaffStatus) (*entities.StaffStatus, error) {
	status := models.StaffStatus{}
	status.FromEntity(req)

	if err := r.db.Create(&status).Error; err != nil {
		return nil, fmt.Errorf("error create status of staff: %w", err)
	}

	result := status.ToEntity()
	return result, nil
}

func (r *StaffStatusRepositoryImpl) RemoveStaffStatus(id uint) error {
	status := models.StaffStatus{}

	if err := r.db.Where("id = ?", id).Delete(&status).Error; err != nil {
		return fmt.Errorf("error deleting status of staff %w", err)
	}

	return nil
}

func (r *StaffStatusRepositoryImpl) GetUserStatus(staffStatusID uint) (*entities.UserStatus, error) {
	statusMapping := models.StatusMapping{}

	if err := r.db.Where("staff_status_id = ?", staffStatusID).First(&statusMapping).Error; err != nil {
		return nil, fmt.Errorf("error finding mapping: %w", err)
	}

	userStatus := models.UserStatus{}
	if err := r.db.First(&userStatus, statusMapping.UserStatusID).Error; err != nil {
		return nil, fmt.Errorf("error finding user status: %w", err)
	}

	result := userStatus.ToEntity()

	return result, nil
}

func (r *StaffStatusRepositoryImpl) BindStatusMapping(mappings *[]entities.StatusMapping) error {
	for _, m := range *mappings {
		mapping := models.StatusMapping{}
		err := r.db.Where("staff_status_id = ?", m.StaffStatusID).First(&mapping).Error

		if err == nil {
			if mapping.UserStatusID != m.UserStatusID {
				mapping.UserStatusID = m.UserStatusID
				if err := r.db.Save(&mapping).Error; err != nil {
					return fmt.Errorf("error updating mapping: %w", err)
				}
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			newMapping := models.StatusMapping{
				StaffStatusID: m.StaffStatusID,
				UserStatusID:  m.UserStatusID,
			}
			if err := r.db.Create(&newMapping).Error; err != nil {
				return fmt.Errorf("error creating new mapping: %w", err)
			}
		} else {
			return fmt.Errorf("error querying mapping: %w", err)
		}
	}
	return nil
}

func (r *StaffStatusRepositoryImpl) GetStaffStatus() (*[]entities.StaffStatus, error) {
	status := []models.StaffStatus{}

	if err := r.db.Find(&status).Error; err != nil {
		return nil, fmt.Errorf("error finding staffstatus: %w", err)
	}

	entityStaffStatus := make([]entities.StaffStatus, len(status))
	for i, s := range status {
		entityStaffStatus[i] = *s.ToEntity()
	}
	return &entityStaffStatus, nil
}
