package interfaces

import (
	"orders/internal/domain"

	"github.com/google/uuid"
)

type OrderRepository interface {
	CreateOrder(order *domain.Order) error
	GetOrder(orderID uuid.UUID) (*domain.Order, error)
	GetOrders() ([]domain.Order, error)
	UpdateOrder(order *domain.Order) error
	DeleteOrder(orderID uuid.UUID) error
}
