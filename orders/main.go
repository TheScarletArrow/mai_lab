package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"

	"time"
)

type Order struct {
	Id         int      `json:"id"`
	CustomerId int      `json:"customerId"`
	Items      []string `json:"items"`
	Total      float64  `json:"total"`
}

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	createOrder(*c, Order{1, 1, []string{"item1", "item2"}, 1100.00})
	createOrder(*c, Order{2, 2, []string{"item3", "item4"}, 200.00})

	//measure time
	start := time.Now()
	fmt.Println(getOrder(*c, 1))
	end := time.Now()
	fmt.Printf("%v", end.Sub(start))
}

//create function that returns Order and takes arguments of customerId, Items and Total

func createOrder(cache cache.Cache, order Order) Order {

	cache.Set("order_"+(string(order.Id)), order, 5*time.Minute)
	return order
}

// read order, get from cache

func getOrder(cache cache.Cache, id int) Order {
	order, found := cache.Get("order_" + (string(id)))
	if found {
		return order.(Order)
	}
	return Order{}
}

//delete Order by id

func deleteOrder(cache cache.Cache, id int) {
	cache.Delete("order_" + (string(id)))
}
