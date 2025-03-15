package catalog

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"ordering_service/pb/catalog"
	"time"
)

type RestaurantServiceClient struct {
	RestaurantServiceClient catalog.RestaurantServiceClient
	Context                 context.Context
}

func NewRestaurantServiceClient() (RestaurantServiceClientInterface, *grpc.ClientConn) {
	DELIVERY_SERVICE_URL := "localhost:8081"
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	connection, err := grpc.NewClient(DELIVERY_SERVICE_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := catalog.NewRestaurantServiceClient(connection)

	return &RestaurantServiceClient{
		RestaurantServiceClient: client,
		Context:                 ctx,
	}, connection
}

func (r *RestaurantServiceClient) GetRestaurantById(request *catalog.GetRestaurantByIdRequest) (*catalog.Restaurant, error) {
	return r.RestaurantServiceClient.GetRestaurantById(r.Context, request)
}
