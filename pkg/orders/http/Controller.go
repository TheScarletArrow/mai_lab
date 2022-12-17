package http

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"orders/orders/models"
	"orders/orders/repository"
)

type OrderController interface {
	GetOrderById(id string) (models.Order, error)
}

func CreateOrder(order models.Order) (models.Order, error) {
	insert := `INSERT INTO orders (id, customer_id, items, total) VALUES ($1, $2, $3, $4)`
	db := sql.DB{}
	db.Exec(insert, order.Id, order.CustomerId, order.Items, order.Total)
	return repository.CreateOrder(order), nil

}

//function to handle http requests

func HandleRequests() {
	//get body of request
	//handle POST request

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		//get body of request
		var order models.Order
		_ = json.NewDecoder(r.Body).Decode(&order)
		order.Id = uuid.New().String()
		//create order
		createdOrder, _ := CreateOrder(order)
		//return order
		json.NewEncoder(w).Encode(createdOrder)

	})

	http.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		//get id from request
		id := r.URL.Path[len("/orders/"):]
		//get order by id
		order, _ := GetOrderById(id)
		//return order
		json.NewEncoder(w).Encode(order)

	})

}

func GetOrderById(id string) (models.Order, error) {

	return repository.GetOrderById(id)

}
