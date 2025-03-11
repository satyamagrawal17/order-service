package service

import (
	"ordering_service/internal/client"
	"ordering_service/internal/dto"
	"ordering_service/internal/model"
	"ordering_service/internal/repository"
	"ordering_service/pb/delivery"
)

type OrderService struct {
	repository      repository.IOrderRepository
	deliveryService client.DeliveryServiceClientInterface
}

func NewOrderService(orderRepository repository.IOrderRepository, deliveryService client.DeliveryServiceClientInterface) *OrderService {
	return &OrderService{
		repository:      orderRepository,
		deliveryService: deliveryService,
	}
}

func (orderService *OrderService) CreateOrder(createOrderQuery dto.CreateOrderRequest) *model.Order {

	if len(createOrderQuery.MenuItemList) == 0 {
		panic("Atleast one menu item is required to place order")
	}

	order := &model.Order{
		UserId: createOrderQuery.UserId,
		Status: "PENDING",
	}

	menuItemsToCreate := []model.OrderItem{}
	for _, menuItem := range createOrderQuery.MenuItemList {
		menuItemsToCreate = append(menuItemsToCreate, model.OrderItem{
			MenuItemId:   menuItem.MenuItemId,
			RestaurantId: menuItem.RestaurantId,
			Quantity:     menuItem.Quantity,
			OrderId:      order.ID,
		})
	}

	order.OrderItems = menuItemsToCreate
	orderService.repository.Create(order)

	orderService.deliveryService.AssignDelivery(&delivery.AssignDeliveryRequest{
		OrderId:      order.ID,
		RestaurantId: createOrderQuery.MenuItemList[0].RestaurantId,
	})

	return order
}

func (orderService *OrderService) GetOrderById(orderId uint32) *model.Order {
	return orderService.repository.GetOrderById(orderId)
}

func (orderService *OrderService) GetOrdersByUserId(userId uint32) []model.Order {
	return orderService.repository.GetOrdersByUserId(userId)
}

func (orderService *OrderService) UpdateStatus(orderId uint32, status string) *model.Order {
	return orderService.repository.UpdateStatus(orderId, status)
}

func (orderService *OrderService) GetOrders() []model.Order {
	return orderService.repository.GetOrders()
}
