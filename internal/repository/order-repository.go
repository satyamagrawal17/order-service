package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"ordering_service/internal/database"
)

type OrderRepository struct {
	DB                 *dynamodb.Client
	OrderTableName     string
	OrderItemTableName string
}

func NewOrderRepository(dy *database.DynamoDB) IOrderRepository {
	return &OrderRepository{
		DB:                 dy.DB,
		OrderTableName:     dy.OrderTableName,
		OrderItemTableName: dy.OrderItemTableName,
	}
}
