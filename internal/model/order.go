package model

import "time"

//type Order struct {
//	Id         string      `dynamodbav:"id"`
//	UserId     string      `dynamodbav:"user_id"`
//	Status     string      `dynamodbav:"status"`
//	CreatedAt  string      `dynamodbav:"created_at"`
//	UpdatedAt  string      `dynamodbav:"updated_at"`
//	OrderItems []OrderItem `dynamodbav:"order_items"`
//}

type Order struct {
	ID         uint32 `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserId     uint32
	Status     string
	CreatedAt  time.Time   `gorm:"autoCreateTime:false"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime:false"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`
}
