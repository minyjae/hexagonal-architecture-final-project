package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type OrderRepository interface {
	UpdateOrder(orderMapping *entities.OrderMapping) (*entities.Order, error)
	CreateOrder(req *entities.Order) (*entities.Order, error)
	RemoveOrder(orderID uint) error
}
