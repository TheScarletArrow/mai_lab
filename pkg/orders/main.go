package main

import (
	"database/sql"
	"fmt"
	"orders/orders/http"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"orders/orders/models"
	"orders/orders/repository"
	//"time"
)

func main() {

	// create table orders

	create := `CREATE TABLE orders (
    		id string PRIMARY KEY,
    		customer_id int,
    		items []string,
    		total float64
    		    		)`
	db := sql.DB{}
	db.Exec(create)

	c := cache.New(10*time.Minute, 5*time.Minute)
	defer c.Flush()
	//create server on port 8081

	createOrder, err := http.CreateOrder(models.Order{Id: uuid.New().String(), CustomerId: 1, Items: []string{}, Total: 10.0})
	if err != nil {

	}
	fmt.Println(createOrder)
	order1 := repository.CreateOrder(models.Order{Id: uuid.New().String(), CustomerId: 1, Items: []string{"item1", "item2"}, Total: 1100.00})
	fmt.Println(order1)
	//order2 := repository.CreateOrder(*c, models.Order{Id: uuid.New(), CustomerId: 2, Items: []string{"item3", "item4"}, Total: 200.00})

	//measure time
	//start := time.Now()
	//fmt.Println(repository.GetOrderById(*c, order1.Id))
	//end := time.Now()
	//fmt.Printf("%v", end.Sub(start))
	////deleteOrder(*c, order2.Id)
	//fmt.Println(repository.GetOrderById(*c, uuid.New().String()))

}

//create function that returns Order and takes arguments of customerId, Items and Total

//func createOrder(cache cache.Cache, order models.Order) models.Order {
//
//	cache.Set("order_"+(order.Id.String()), order, 5*time.Minute)
//	return order
//}
//
//// read order, get from cache
//
//func getOrder(cache cache.Cache, id uuid.UUID) models.Order {
//	order, found := cache.Get("order_" + id.String())
//	if found {
//		return order.(models.Order)
//	}
//	return models.Order{}
//}
//
////delete Order by id
//
//func deleteOrder(cache cache.Cache, id uuid.UUID) {
//	cache.Delete("order_" + id.String())
//}
