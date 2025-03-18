package catalog

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"ordering_service/pb/catalog"
	"time"
)

type ItemServiceClient struct {
	ItemServiceClient catalog.ItemServiceClient
	Context           context.Context
}

func NewItemServiceClient() (ItemServiceClientInterface, *grpc.ClientConn) {
	DELIVERY_SERVICE_URL := "localhost:8081"
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	connection, err := grpc.NewClient(DELIVERY_SERVICE_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := catalog.NewItemServiceClient(connection)

	return &ItemServiceClient{
		ItemServiceClient: client,
		Context:           ctx,
	}, connection
}

func (i *ItemServiceClient) GetItemById(request *catalog.GetItemByIdRequest) (*catalog.MenuItem, error) {
	return i.ItemServiceClient.GetItemById(i.Context, request)
}
