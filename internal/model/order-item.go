package model

type OrderItem struct {
	Id           string `dynamodbav:"id"`
	OrderId      string `dynamodbav:"order_id"` // Store OrderID as string
	MenuItemId   string `dynamodbav:"menu_item_id"`
	RestaurantId string `dynamodbav:"restaurant_id"`
	Quantity     int    `dynamodbav:"quantity"`
	CreatedAt    string `dynamodbav:"created_at"`
	UpdatedAt    string `dynamodbav:"updated_at"`
}
