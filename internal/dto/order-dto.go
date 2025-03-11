package dto

type MenuItemRequest struct {
	MenuItemId   uint32
	RestaurantId uint32
	Quantity     int
}

type CreateOrderRequest struct {
	UserId       uint32
	MenuItemList []MenuItemRequest
}

type UpdateOrderRequest struct {
	OrderId uint32
	Status  string
}
