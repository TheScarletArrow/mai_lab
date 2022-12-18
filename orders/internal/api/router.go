package api

import (
	"github.com/gorilla/mux"
)

func NewRouter(orderHandler *OrderHandler) *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	muxRouter.HandleFunc("/orders", orderHandler.GetOrders).Methods("GET")
	muxRouter.HandleFunc("/orders/{order_id}", orderHandler.GetOrder).Methods("GET")
	muxRouter.HandleFunc("/orders/{order_id}", orderHandler.UpdateOrder).Methods("PUT")
	muxRouter.HandleFunc("/orders/{order_id}", orderHandler.DeleteOrder).Methods("DELETE")
	return muxRouter
}
