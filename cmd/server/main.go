package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"ordering_service/internal/client/catalog"
	"ordering_service/internal/client/delivery"
	"ordering_service/internal/database"
	"ordering_service/internal/repository"
	"ordering_service/internal/server"
	"ordering_service/internal/service"
	"ordering_service/pb/order"
)

func Start() {
	dbEngine, err := database.InitDatabaseEngine()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	orderRepository := repository.NewOrderRepository(dbEngine)
	deliveryService, connection := delivery.NewDeliveryServiceClient()
	restaurantService, restConnection := catalog.NewRestaurantServiceClient()
	itemService, itemConnection := catalog.NewItemServiceClient()
	defer connection.Close()
	defer restConnection.Close()
	defer itemConnection.Close()
	orderService := service.NewOrderService(orderRepository, deliveryService, restaurantService, itemService)
	grpcServer := server.NewGRPCServer(orderService)

	s := grpc.NewServer()
	order.RegisterOrderServiceServer(s, grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
