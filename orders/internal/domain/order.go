package domain

import (
	"github.com/google/uuid"
)

type Order struct {
	OrderID uuid.UUID `json:"order_id" db:"order_id"`
	UserID  uuid.UUID `json:"user_id" db:"user_id"`
	Price   float64   `json:"price" db:"price"`
}
