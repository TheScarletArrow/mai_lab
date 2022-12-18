package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"orders/internal/api"
	"orders/internal/infrastructure"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx := context.Background()
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	log.Default().Println("Connected to MongoDB!")

	ordersRepository := infrastructure.NewMongoOrderRepository(client.Database("orders"))
	orderHandler := api.NewOrderHandler(ordersRepository)
	router := api.NewRouter(orderHandler)
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
