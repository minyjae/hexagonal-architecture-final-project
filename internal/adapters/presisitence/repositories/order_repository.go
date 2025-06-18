package repositories

import (
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presisitence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) CreateOrder(req *entities.Order) (*entities.Order, error) {
	order := &models.Order{}
	order.FromEntity(req)

	if err := r.db.Create(&order).Error; err != nil {
		return nil, fmt.Errorf("failed to create new order: %w", err)
	}

	r.addNewOrderToListQueue(order.ToEntity())

	createdOrder := models.Order{}
	if err := r.db.Preload("OrderMappings").First(&createdOrder, order.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to preload created order: %w", err)
	}

	result := createdOrder.ToEntity()
	return result, nil
}

func (r *OrderRepositoryImpl) UpdateOrder(orderMapping *entities.OrderMapping) (*entities.Order, error) {
	findOrder := &models.OrderMapping{}

	if err := r.db.Where("order_id = ? AND list_queue_id = ?", orderMapping.OrderID, orderMapping.ListQueueID).First(&findOrder).Error; err != nil {
		return nil, fmt.Errorf("failed to find existing order mapping: %w", err)
	}

	findOrder.Checked = orderMapping.Checked

	if err := r.db.Save(&findOrder).Error; err != nil {
		return nil, fmt.Errorf("failed to update order mapping: %w", err)
	}

	updateOrder := models.Order{}
	if err := r.db.Preload("OrderMappings").First(&updateOrder, orderMapping.OrderID).Error; err != nil {
		return nil, fmt.Errorf("failed to preload updated order: %w", err)
	}

	result := updateOrder.ToEntity()
	return result, nil
}

func (r *OrderRepositoryImpl) RemoveOrder(orderID uint) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_id = ?", orderID).Delete(&models.OrderMapping{}).Error; err != nil {
			return fmt.Errorf("failed to delete OrderMapping: %w", err)
		}

		if err := tx.Delete(&models.Order{}, orderID).Error; err != nil {
			return fmt.Errorf("failed to delete Order: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	r.removeOrderFromListQueue(orderID)
	return nil
}

func (r *OrderRepositoryImpl) CreateOrderForListQueue(listQueueID uint) error {
	orders := []models.Order{}

	if err := r.db.Find(&orders).Error; err != nil {
		return fmt.Errorf("failed to fetch orders: %w", err)
	}

	for _, order := range orders {
		orderMapping := models.OrderMapping{
			ListQueueID: listQueueID,
			OrderID:     order.ID,
		}
		if err := r.db.Create(&orderMapping).Error; err != nil {
			return fmt.Errorf("failed to create OrderMapping for orderID %d: %w", order.ID, err)
		}
	}

	return nil
}

func (r *OrderRepositoryImpl) addNewOrderToListQueue(order *entities.Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		listQueue := []models.ListQueue{}

		if err := tx.Where("users_status_id IN ?", []int{1, 2}).Find(&listQueue).Error; err != nil {
			return err
		}

		orderMapping := []models.OrderMapping{}
		for _, list := range listQueue {
			orderMapping = append(orderMapping, models.OrderMapping{
				ListQueueID: list.ID,
				OrderID:     order.ID,
				Checked:     false,
			})
		}

		if err := tx.Create(&orderMapping).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *OrderRepositoryImpl) removeOrderFromListQueue(orderID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_id = ?", orderID).Delete(&models.OrderMapping{}).Error; err != nil {
			return fmt.Errorf("failed to delete OrderMapping: %w", err)
		}

		if err := tx.Delete(&models.Order{}, orderID).Error; err != nil {
			return fmt.Errorf("failed to delete Order: %w", err)
		}

		return nil
	})
}
