package repositories

import (
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presisitence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type UserStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewUserStatusRepository(db *gorm.DB) *UserStatusRepositoryImpl {
	return &UserStatusRepositoryImpl{db: db}
}

func (r *UserStatusRepositoryImpl) CreateUserStatus(req *entities.UserStatus) (*entities.UserStatus, error) {
	status := models.UserStatus{}
	status.FromEntity(req)

	if err := r.db.Create(&status).Error; err != nil {
		return nil, fmt.Errorf("error creating status of user: %w", err)
	}

	result := status.ToEntity()

	return result, nil
}

func (r *UserStatusRepositoryImpl) RemoveUserStatus(id uint) error {
	status := models.UserStatus{}

	if err := r.db.Where("id = ?", id).Delete(&status).Error; err != nil {
		return fmt.Errorf("error deleting status of user: %w", err)
	}

	return nil
}

func (r *UserStatusRepositoryImpl) GetUserStatus() (*[]entities.UserStatus, error) {
	status := []models.UserStatus{}

	if err := r.db.Find(&status).Error; err != nil {
		return nil, fmt.Errorf("error finding userstatus")
	}

	entityUserStatus := make([]entities.UserStatus, len(status))
	for i, e := range status {
		entityUserStatus[i] = *e.ToEntity()
	}

	return &entityUserStatus, nil
}
