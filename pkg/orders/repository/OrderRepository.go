package repository

import (
	"errors"
	"github.com/google/uuid"
	"orders/orders/models"
)

// interface
type OrderRepository interface {
	GetOrderById(id uuid.UUID) (models.Order, error)

	CreateOrder(order *models.Order) models.Order

	DeleteOrder(id uuid.UUID) error
	NewOrderRepository() interface{}
}

// implement GetOrderById
func GetOrderById(id string) (models.Order, error) {
	//order, found := cache.Get("order_" + id)
	//if found {
	//	return order.(models.Order), nil
	//}
	return models.Order{}, errors.New("Order not found")
}

func CreateOrder(order models.Order) models.Order {

	//cache.Set("order_"+(order.Id), order, 5*time.Minute)
	return order
}
