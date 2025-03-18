package model

import "time"

//type OrderItem struct {
//	Id           string `dynamodbav:"id"`
//	OrderId      string `dynamodbav:"order_id"` // Store OrderID as string
//	MenuItemId   string `dynamodbav:"menu_item_id"`
//	RestaurantId string `dynamodbav:"restaurant_id"`
//	Quantity     int    `dynamodbav:"quantity"`
//	CreatedAt    string `dynamodbav:"created_at"`
//	UpdatedAt    string `dynamodbav:"updated_at"`
//}

type OrderItem struct {
	ID           uint32 `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	OrderId      uint32 `gorm:"foreignKey:OrderId"`
	MenuItemId   uint32
	RestaurantId uint32
	Quantity     int
	CreatedAt    time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:true"`
}
