package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"orders/internal/domain"
	"orders/internal/interfaces"
)

type OrderHandler struct {
	orderRepository interfaces.OrderRepository
}

func NewOrderHandler(orderRepository interfaces.OrderRepository) *OrderHandler {
	return &OrderHandler{
		orderRepository: orderRepository,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// search by OrderId
	getOrder, err := h.orderRepository.GetOrder(order.OrderID)
	if getOrder != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Order already exists with this ID " + order.OrderID.String()))
		return
	}

	err = h.orderRepository.CreateOrder(&order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := uuid.Parse(mux.Vars(r)["order_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := h.orderRepository.GetOrder(orderID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderRepository.GetOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.orderRepository.UpdateOrder(&order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {

	orderID, err := uuid.Parse(mux.Vars(r)["order_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	getOrder, err := h.orderRepository.GetOrder(orderID)
	if getOrder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Order does not exist with ID " + orderID.String()))
		return
	}
	err = h.orderRepository.DeleteOrder(orderID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
