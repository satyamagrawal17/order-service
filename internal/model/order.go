package model

type Order struct {
	Id         string      `dynamodbav:"id"`
	UserId     string      `dynamodbav:"user_id"`
	Status     string      `dynamodbav:"status"`
	CreatedAt  string      `dynamodbav:"created_at"`
	UpdatedAt  string      `dynamodbav:"updated_at"`
	OrderItems []OrderItem `dynamodbav:"order_items"`
}
