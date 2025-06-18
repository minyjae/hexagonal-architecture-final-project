package repositories

import (
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presisitence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type ListQueueRepositoryImpl struct {
	db *gorm.DB
}

func NewListQueueRepository(db *gorm.DB) *ListQueueRepositoryImpl {
	return &ListQueueRepositoryImpl{db: db}
}

func (r *ListQueueRepositoryImpl) CreateListQueue(req *entities.ListQueue) (*entities.ListQueue, error) {
	// ให้ listQueue เป็นตัวแปลที่มีข้อมูลของ models.ListQueue
	listQueue := &models.ListQueue{}
	// ให้ listQueue ไปอ่านข้อมูลของ req ให้เป็น models เพื่อจะได้คุยกับ database ผ่าน gorm ได้
	listQueue.FromEntity(req)

	if err := r.db.Create(listQueue).Error; err != nil {
		return nil, err
	}

	result := &models.ListQueue{}
	err := r.db.Preload("StaffStatus").Preload("UserStatus").First(result, req.ID).Error
	if err != nil {
		return nil, fmt.Errorf("error to preload relation: %w", err)
	}
	// แปลงกลับออกมาเป็น entity ที่จะนำไปใช้ต่อได้
	*req = *listQueue.ToEntity()
	return req, nil
}

func (r *ListQueueRepositoryImpl) UpdateListQueue(req *entities.ListQueue) (*entities.ListQueue, error) {
	listQueue := &models.ListQueue{}
	listQueue.FromEntity(req)

	err := r.db.Model(listQueue).Where("priority = ?", req.Priority).Updates(req).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update list queue: %w", err)
	}

	result := &models.ListQueue{}
	err = r.db.Preload("StaffStatus").Preload("UserStatus").First(result, req.ID).Error
	if err != nil {
		return nil, fmt.Errorf("error to preload relation: %w", err)
	}

	entityResult := result.ToEntity()
	return entityResult, nil
}

func (r *ListQueueRepositoryImpl) GetListQueue() (*[]entities.ListQueue, error) {
	listQueues := []models.ListQueue{}

	err := r.db.Find(&listQueues).Error
	if err != nil {
		return nil, fmt.Errorf("error finding list queue: %w", err)
	}

	entityListQueues := make([]entities.ListQueue, len(listQueues))
	for i, model := range listQueues {
		entityListQueues[i] = *model.ToEntity()
	}

	return &entityListQueues, nil
}

func (r *ListQueueRepositoryImpl) GetListQueueByStaffStatus(statusID uint) (*[]entities.ListQueue, error) {
	listQueues := []models.ListQueue{}

	err := r.db.Where("staff_status_id = ?", statusID).Find(&listQueues).Error
	if err != nil {
		return nil, fmt.Errorf("error finding list queue by staff status: %w", err)
	}

	entityListQueues := make([]entities.ListQueue, len(listQueues))
	for i, model := range listQueues {
		entityListQueues[i] = *model.ToEntity()
	}

	return &entityListQueues, nil
}

func (r *ListQueueRepositoryImpl) GetListQueueByUserStatus(statusID uint) (*[]entities.ListQueue, error) {
	listQueues := []models.ListQueue{}

	err := r.db.Where("user_status_id = ?", statusID).Find(&listQueues).Error
	if err != nil {
		return nil, fmt.Errorf("error finding list queue by user status: %w", err)
	}

	entityListQueue := make([]entities.ListQueue, len(listQueues))
	for i, model := range listQueues {
		entityListQueue[i] = *model.ToEntity()
	}

	return &entityListQueue, nil
}

func (r *ListQueueRepositoryImpl) GetListQueueByID(id uint) (*entities.ListQueue, error) {
	listQueue := &models.ListQueue{}

	err := r.db.Where("id = ?", id).First(&listQueue).Error
	if err != nil {
		return nil, fmt.Errorf("error finding list queue by id: %w", err)
	}

	resultListQueue := *listQueue.ToEntity()
	return &resultListQueue, nil
}

func (r *ListQueueRepositoryImpl) UpdateStaffStatusListQueue(id, staffID uint) (*entities.ListQueue, error) {
	listQueue := &models.ListQueue{}

	err := r.db.Where("id = ?", id).First(&listQueue).Error
	if err != nil {
		return nil, fmt.Errorf("error finding list queue: %w", err)
	}

	statusMapping := &models.StatusMapping{}
	err = r.db.Where("staff_status_id = ?", staffID).First(&statusMapping).Error
	if err != nil {
		return nil, fmt.Errorf("error finding status mapping: %w", err)
	}

	listQueue.StaffStatusID = staffID
	listQueue.UsersStatusID = statusMapping.UserStatusID

	if err := r.db.Save(&listQueue).Error; err != nil {
		return nil, fmt.Errorf("failed to update staff status of list queue: %w", err)
	}

	preloadStatus := &models.ListQueue{}
	err = r.db.Preload("StaffStatus").Preload("UserStatus").First(&preloadStatus, listQueue.ID).Error
	if err != nil {
		return nil, fmt.Errorf("error to preload relation: %w", err)
	}

	result := *preloadStatus.ToEntity()

	return &result, nil
}

func (r *ListQueueRepositoryImpl) UpdatePriority(id uint, priority int) (*entities.ListQueue, error) {
	listQueue := models.ListQueue{}

	err := r.db.Where("id = ?", id).First(&listQueue).Error
	if err != nil {
		return nil, fmt.Errorf("error finding list queue: %w", err)
	}

	if listQueue.UsersStatusID == 3 {
		listQueue.Priority = 0
	} else {
		listQueue.Priority = priority
	}

	if err = r.db.Save(&listQueue).Error; err != nil {
		return nil, fmt.Errorf("failed to update priority of list queue: %w", err)
	}

	result := *listQueue.ToEntity()

	return &result, nil
}
