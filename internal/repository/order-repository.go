package repository

import (
	"ordering_service/internal/database"
	"ordering_service/internal/model"
)

type OrderRepository struct {
	databaseEngine *database.DatabaseEngine
}

func NewOrderRepository(databaseEngine *database.DatabaseEngine) IOrderRepository {
	return &OrderRepository{
		databaseEngine: databaseEngine,
	}
}

func (orderRepository *OrderRepository) Create(
	order *model.Order,
) {
	orderRepository.databaseEngine.Db.Create(order)
}

func (orderRepository *OrderRepository) UpdateStatus(orderId uint32, status string) *model.Order {
	order := orderRepository.GetOrderById(orderId)
	order.Status = status
	orderRepository.databaseEngine.Db.Save(order)
	return order
}

func (orderRepository *OrderRepository) GetOrderById(orderId uint32) *model.Order {
	var order model.Order
	orderRepository.databaseEngine.Db.Preload("OrderItems").Find(&order).Where(&model.Order{ID: orderId})
	return &order
}

func (orderRepository *OrderRepository) GetOrders() []model.Order {
	var orders []model.Order
	orderRepository.databaseEngine.Db.Find(&orders)
	return orders
}

func (orderRepository *OrderRepository) GetOrdersByUserId(userId uint32) []model.Order {
	var orders []model.Order
	orderRepository.databaseEngine.Db.Find(orders).Where(&model.Order{UserId: userId})
	return orders
}
