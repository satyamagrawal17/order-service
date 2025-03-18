package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	catalog2 "ordering_service/internal/client/catalog"
	delivery2 "ordering_service/internal/client/delivery"
	"ordering_service/internal/dto"
	"ordering_service/internal/model"
	"ordering_service/internal/repository"
	"ordering_service/pb/catalog"
	"ordering_service/pb/delivery"
)

func TestOrderService_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := repository.NewMockIOrderRepository(ctrl)
	mockDeliveryService := delivery2.NewMockDeliveryServiceClientInterface(ctrl)
	mockRestaurantService := catalog2.NewMockRestaurantServiceClientInterface(ctrl)
	mockItemService := catalog2.NewMockItemServiceClientInterface(ctrl)

	orderService := NewOrderService(mockOrderRepo, mockDeliveryService, mockRestaurantService, mockItemService)

	createOrderRequest := dto.CreateOrderRequest{
		UserId: 1,
		MenuItemList: []dto.MenuItemRequest{
			{MenuItemId: 1, RestaurantId: 1, Quantity: 2},
		},
	}

	mockRestaurantService.EXPECT().GetRestaurantById(gomock.Any()).Return(&catalog.Restaurant{}, nil).Times(1)
	mockItemService.EXPECT().GetItemById(gomock.Any()).Return(&catalog.MenuItem{}, nil).Times(1)
	mockOrderRepo.EXPECT().Create(gomock.Any()).Times(1)
	mockDeliveryService.EXPECT().AssignDelivery(gomock.Any()).Return(&delivery.AssignDeliveryResponse{}, nil).Times(1)

	order, _ := orderService.CreateOrder(createOrderRequest)
	if order.UserId != createOrderRequest.UserId {
		t.Errorf("expected userId %v, got %v", createOrderRequest.UserId, order.UserId)
	}
	if order.Status != "PENDING" {
		t.Errorf("expected status PENDING, got %v", order.Status)
	}
	if len(order.OrderItems) != len(createOrderRequest.MenuItemList) {
		t.Errorf("expected %v order items, got %v", len(createOrderRequest.MenuItemList), len(order.OrderItems))
	}
}

func TestOrderService_CreateOrder_EmptyMenuItemList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := repository.NewMockIOrderRepository(ctrl)
	mockDeliveryService := delivery2.NewMockDeliveryServiceClientInterface(ctrl)
	mockRestaurantService := catalog2.NewMockRestaurantServiceClientInterface(ctrl)
	mockItemService := catalog2.NewMockItemServiceClientInterface(ctrl)

	orderService := NewOrderService(mockOrderRepo, mockDeliveryService, mockRestaurantService, mockItemService)

	createOrderRequest := dto.CreateOrderRequest{
		UserId:       1,
		MenuItemList: []dto.MenuItemRequest{},
	}

	_, err := orderService.CreateOrder(createOrderRequest)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}

func TestOrderService_GetOrderById_NonExistentOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := repository.NewMockIOrderRepository(ctrl)
	mockDeliveryService := delivery2.NewMockDeliveryServiceClientInterface(ctrl)
	mockRestaurantService := catalog2.NewMockRestaurantServiceClientInterface(ctrl)
	mockItemService := catalog2.NewMockItemServiceClientInterface(ctrl)

	orderService := NewOrderService(mockOrderRepo, mockDeliveryService, mockRestaurantService, mockItemService)

	mockOrderRepo.EXPECT().GetOrderById(uint32(1)).Return(nil).Times(1)

	order := orderService.GetOrderById(1)
	if order != nil {
		t.Errorf("expected nil, got %v", order)
	}
}

func TestOrderService_UpdateStatus_ValidData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := repository.NewMockIOrderRepository(ctrl)
	mockDeliveryService := delivery2.NewMockDeliveryServiceClientInterface(ctrl)
	mockRestaurantService := catalog2.NewMockRestaurantServiceClientInterface(ctrl)
	mockItemService := catalog2.NewMockItemServiceClientInterface(ctrl)

	orderService := NewOrderService(mockOrderRepo, mockDeliveryService, mockRestaurantService, mockItemService)

	order := &model.Order{ID: 1, Status: "COMPLETED"}
	mockOrderRepo.EXPECT().UpdateStatus(uint32(1), "COMPLETED").Return(order).Times(1)

	updatedOrder := orderService.UpdateStatus(1, "COMPLETED")
	if updatedOrder.Status != "COMPLETED" {
		t.Errorf("expected status COMPLETED, got %v", updatedOrder.Status)
	}
}

func TestOrderService_GetOrdersByUserId_NoOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := repository.NewMockIOrderRepository(ctrl)
	mockDeliveryService := delivery2.NewMockDeliveryServiceClientInterface(ctrl)
	mockRestaurantService := catalog2.NewMockRestaurantServiceClientInterface(ctrl)
	mockItemService := catalog2.NewMockItemServiceClientInterface(ctrl)

	orderService := NewOrderService(mockOrderRepo, mockDeliveryService, mockRestaurantService, mockItemService)

	mockOrderRepo.EXPECT().GetOrdersByUserId(uint32(1)).Return([]model.Order{}).Times(1)

	orders := orderService.GetOrdersByUserId(1)
	if len(orders) != 0 {
		t.Errorf("expected 0 orders, got %v", len(orders))
	}
}
