package service

import "ordering_service/internal/repository"

type OrderService struct {
	repository repository.IOrderRepository
}

func NewOrderService(orderRepository repository.IOrderRepository) *OrderService {
	return &OrderService{
		repository: orderRepository,
	}
}
