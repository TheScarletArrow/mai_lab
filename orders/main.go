package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"time"
)

type Order struct {
	Id         uuid.UUID `json:"id"`
	CustomerId int       `json:"customerId"`
	Items      []string  `json:"items"`
	Total      float64   `json:"total"`
}

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	order1 := createOrder(*c, Order{uuid.New(), 1, []string{"item1", "item2"}, 1100.00})
	order2 := createOrder(*c, Order{uuid.New(), 2, []string{"item3", "item4"}, 200.00})

	//measure time
	start := time.Now()
	fmt.Println(getOrder(*c, order1.Id))
	end := time.Now()
	fmt.Printf("%v", end.Sub(start))
	deleteOrder(*c, order2.Id)
	fmt.Println(getOrder(*c, order2.Id))

}

//create function that returns Order and takes arguments of customerId, Items and Total

func createOrder(cache cache.Cache, order Order) Order {

	cache.Set("order_"+(order.Id.String()), order, 5*time.Minute)
	return order
}

// read order, get from cache

func getOrder(cache cache.Cache, id uuid.UUID) Order {
	order, found := cache.Get("order_" + id.String())
	if found {
		return order.(Order)
	}
	return Order{}
}

//delete Order by id

func deleteOrder(cache cache.Cache, id uuid.UUID) {
	cache.Delete("order_" + id.String())
}
