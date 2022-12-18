package infrastructure

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"orders/internal/domain"
	"orders/internal/interfaces"
)

const (
	ordersCollection string = "orders"
	timeout                 = 10 * time.Second
)

type mongoOrderRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) interfaces.OrderRepository {
	return &mongoOrderRepository{
		collection: db.Collection(ordersCollection),
	}
}

func (m *mongoOrderRepository) CreateOrder(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := m.collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoOrderRepository) GetOrder(orderID uuid.UUID) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var order domain.Order
	err := m.collection.FindOne(ctx, bson.M{"orderid": orderID}).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (m *mongoOrderRepository) GetOrders() ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var orders []domain.Order
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err

	}
	return orders, nil
}

func (m *mongoOrderRepository) UpdateOrder(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := m.collection.UpdateOne(ctx, bson.M{"orderid": order.OrderID}, bson.M{"$set": order})
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoOrderRepository) DeleteOrder(orderID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := m.collection.DeleteOne(ctx, bson.M{"orderid": orderID})
	if err != nil {
		return err
	}

	return nil
}
