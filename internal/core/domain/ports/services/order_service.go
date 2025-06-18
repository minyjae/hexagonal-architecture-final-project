package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type OrderService interface {
	CreateOrderForListQueue(listQueueID uint) error
	AddNewOrderToListQueue(order *entities.Order) error
	RemoveOrderFromListQueue(orderID uint) error
}
