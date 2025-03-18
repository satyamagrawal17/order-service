package delivery

import (
	"context"
	"log"
	"ordering_service/pb/delivery"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DeliveryServiceClient struct {
	DeliveryServiceClient delivery.DeliveryServiceClient
	Context               context.Context
}

func NewDeliveryServiceClient() (DeliveryServiceClientInterface, *grpc.ClientConn) {
	DELIVERY_SERVICE_URL := "localhost:50052"
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	connection, err := grpc.NewClient(DELIVERY_SERVICE_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := delivery.NewDeliveryServiceClient(connection)

	return &DeliveryServiceClient{
		DeliveryServiceClient: client,
		Context:               ctx,
	}, connection
}

func (d *DeliveryServiceClient) AssignDelivery(assignDeliveryRequest *delivery.AssignDeliveryRequest) (*delivery.AssignDeliveryResponse, error) {
	return d.DeliveryServiceClient.AssignDelivery(d.Context, assignDeliveryRequest)
}

func (d *DeliveryServiceClient) UpdateDeliveryStatus(request *delivery.UpdateDeliveryStatusRequest) (*delivery.UpdateDeliveryStatusResponse, error) {
	return d.DeliveryServiceClient.UpdateDeliveryStatus(d.Context, request)
}
