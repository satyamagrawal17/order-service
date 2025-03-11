package repository

import "ordering_service/internal/model"

type IOrderRepository interface {
	Create(order *model.Order)
	UpdateStatus(orderId uint32, status string) *model.Order
	GetOrders() []model.Order
	GetOrdersByUserId(userId uint32) []model.Order
	GetOrderById(orderId uint32) *model.Order
}
