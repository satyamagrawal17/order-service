package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"log"
	//"net/http"
	catalog2 "ordering_service/internal/client/catalog"
	delivery2 "ordering_service/internal/client/delivery"
	"ordering_service/internal/dto"
	"ordering_service/internal/model"
	"ordering_service/internal/repository"
	"ordering_service/pb/catalog"
	"ordering_service/pb/delivery"
)

// Define custom error types

type OrderService struct {
	repository        repository.IOrderRepository
	deliveryService   delivery2.DeliveryServiceClientInterface
	restaurantService catalog2.RestaurantServiceClientInterface
	itemService       catalog2.ItemServiceClientInterface
}

func NewOrderService(orderRepository repository.IOrderRepository, deliveryService delivery2.DeliveryServiceClientInterface, restaurantService catalog2.RestaurantServiceClientInterface, itemService catalog2.ItemServiceClientInterface) *OrderService {
	return &OrderService{
		repository:        orderRepository,
		deliveryService:   deliveryService,
		restaurantService: restaurantService,
		itemService:       itemService,
	}
}

func (orderService *OrderService) CreateOrder(createOrderQuery dto.CreateOrderRequest) (*model.Order, error) {

	if len(createOrderQuery.MenuItemList) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "menu item list is empty")
	}

	order := &model.Order{
		UserId: createOrderQuery.UserId,
		Status: "PENDING",
	}
	for _, menuItem := range createOrderQuery.MenuItemList {
		savedRestaurant, err := orderService.restaurantService.GetRestaurantById(&catalog.GetRestaurantByIdRequest{
			Id: uint64(menuItem.RestaurantId),
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get restaurant by id: %v", err)
		}
		if savedRestaurant == nil {
			return nil, status.Errorf(codes.NotFound, "restaurant not found: %v", err)
		}
		savedItem, err := orderService.itemService.GetItemById(&catalog.GetItemByIdRequest{
			MenuItemId: uint64(menuItem.MenuItemId),
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get item by id: %v", err)
		}
		if savedItem == nil {
			return nil, status.Errorf(codes.NotFound, "item not found: %v", err)
		}
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

	return order, nil
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
