package server

import (
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

//func (s *GRPCServer) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
//
//	menuItemList := []dto.MenuItemRequest{}
//	for _, requestOrderItem := range req.OrderItems {
//		menuItemList = append(menuItemList, dto.MenuItemRequest{
//			MenuItemId:   requestOrderItem.MenuItemId,
//			RestaurantId: requestOrderItem.RestaurantId,
//			Quantity:     int(requestOrderItem.Quantity),
//		})
//	}
//
//	createdOrder := s.orderService.CreateOrder(dto.CreateOrderRequest{
//		UserId:       req.UserId,
//		MenuItemList: menuItemList,
//	})
//
//	convertedOrderItems := []*order.OrderItem{}
//	for _, orderItem := range createdOrder.OrderItems {
//		convertedOrderItems = append(convertedOrderItems, &order.OrderItem{
//			MenuItemId:   orderItem.MenuItemId,
//			RestaurantId: orderItem.RestaurantId,
//			Quantity:     int32(orderItem.Quantity),
//		})
//	}
//
//	return &order.CreateOrderResponse{
//		Order: &order.Order{
//			Id:         uint32(createdOrder.ID),
//			UserId:     createdOrder.UserId,
//			Status:     createdOrder.Status,
//			OrderItems: convertedOrderItems,
//		},
//	}, nil
//}
//
//func (s *GRPCServer) GetOrderById(ctx context.Context, req *order.GetOrderRequest) (*order.OrderResponse, error) {
//
//	fetchedOrder := s.orderService.GetOrderById(uint(req.Id))
//
//	convertedOrderItems := []*order.OrderItem{}
//	for _, orderItem := range fetchedOrder.OrderItems {
//		convertedOrderItems = append(convertedOrderItems, &order.OrderItem{
//			MenuItemId:   orderItem.MenuItemId,
//			RestaurantId: orderItem.RestaurantId,
//			Quantity:     int32(orderItem.Quantity),
//		})
//	}
//
//	return &order.OrderResponse{
//		Order: &order.Order{
//			Id:         uint32(fetchedOrder.ID),
//			UserId:     fetchedOrder.UserId,
//			Status:     fetchedOrder.Status,
//			OrderItems: convertedOrderItems,
//		},
//	}, nil
//}
//
//func (s *GRPCServer) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderRequest) (*order.OrderResponse, error) {
//
//	updatedOrder := s.orderService.UpdateStatus(uint(req.OrderId), req.Status)
//
//	return &order.OrderResponse{
//		Order: &order.Order{
//			Id:     uint32(updatedOrder.ID),
//			UserId: updatedOrder.UserId,
//			Status: updatedOrder.Status,
//		},
//	}, nil
//}
//
//func (s *GRPCServer) GetOrderByUserId(ctx context.Context, req *order.GetOrderByUserIdRequest) (*order.MultipleOrderResponse, error) {
//
//	fetchedOrders := s.orderService.GetOrdersByUserId(req.UserId)
//
//	convertedOrderItems := []*order.Order{}
//	for _, orderItem := range fetchedOrders {
//		convertedOrderItems = append(convertedOrderItems, &order.Order{
//			Id:     uint32(orderItem.ID),
//			UserId: orderItem.UserId,
//			Status: orderItem.Status,
//		})
//	}
//
//	return &order.MultipleOrderResponse{
//		Orders: convertedOrderItems,
//	}, nil
//}
