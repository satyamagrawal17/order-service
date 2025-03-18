package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ordering_service/internal/dto"
	"ordering_service/internal/service"
	"ordering_service/pb/order"
)

type GRPCServer struct {
	order.UnimplementedOrderServiceServer
	orderService *service.OrderService
}

func NewGRPCServer(orderService *service.OrderService) *GRPCServer {
	return &GRPCServer{orderService: orderService}
}

func (s *GRPCServer) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	menuItemList := []dto.MenuItemRequest{}
	for _, requestOrderItem := range req.OrderItems {
		menuItemList = append(menuItemList, dto.MenuItemRequest{
			MenuItemId:   requestOrderItem.MenuItemId,
			RestaurantId: requestOrderItem.RestaurantId,
			Quantity:     int(requestOrderItem.Quantity),
		})
	}

	createdOrder, err := s.orderService.CreateOrder(dto.CreateOrderRequest{
		UserId:       req.UserId,
		MenuItemList: menuItemList,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Some error occured: ", err.Error())
	}
	convertedOrderItems := []*order.OrderItem{}
	for _, orderItem := range createdOrder.OrderItems {
		convertedOrderItems = append(convertedOrderItems, &order.OrderItem{
			MenuItemId:   uint32(orderItem.MenuItemId),
			RestaurantId: uint32(orderItem.RestaurantId),
			Quantity:     int32(orderItem.Quantity),
		})
	}

	return &order.CreateOrderResponse{
		Order: &order.Order{
			Id:         uint32(createdOrder.ID),
			UserId:     createdOrder.UserId,
			Status:     createdOrder.Status,
			OrderItems: convertedOrderItems,
		},
	}, nil
}

func (s *GRPCServer) GetOrderById(ctx context.Context, req *order.GetOrderRequest) (*order.OrderResponse, error) {

	fetchedOrder := s.orderService.GetOrderById(req.Id)

	convertedOrderItems := []*order.OrderItem{}
	for _, orderItem := range fetchedOrder.OrderItems {
		convertedOrderItems = append(convertedOrderItems, &order.OrderItem{
			MenuItemId:   orderItem.MenuItemId,
			RestaurantId: orderItem.RestaurantId,
			Quantity:     int32(orderItem.Quantity),
		})
	}

	return &order.OrderResponse{
		Order: &order.Order{
			Id:         uint32(fetchedOrder.ID),
			UserId:     fetchedOrder.UserId,
			Status:     fetchedOrder.Status,
			OrderItems: convertedOrderItems,
		},
	}, nil
}

func (s *GRPCServer) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderRequest) (*order.OrderResponse, error) {

	updatedOrder := s.orderService.UpdateStatus(req.OrderId, req.Status)

	return &order.OrderResponse{
		Order: &order.Order{
			Id:     uint32(updatedOrder.ID),
			UserId: updatedOrder.UserId,
			Status: updatedOrder.Status,
		},
	}, nil
}

func (s *GRPCServer) GetOrderByUserId(ctx context.Context, req *order.GetOrderByUserIdRequest) (*order.MultipleOrderResponse, error) {

	fetchedOrders := s.orderService.GetOrdersByUserId(req.UserId)

	convertedOrderItems := []*order.Order{}
	for _, orderItem := range fetchedOrders {
		convertedOrderItems = append(convertedOrderItems, &order.Order{
			Id:     uint32(orderItem.ID),
			UserId: orderItem.UserId,
			Status: orderItem.Status,
		})
	}

	return &order.MultipleOrderResponse{
		Orders: convertedOrderItems,
	}, nil
}
