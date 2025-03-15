package catalog

import "ordering_service/pb/catalog"

type RestaurantServiceClientInterface interface {
	GetRestaurantById(request *catalog.GetRestaurantByIdRequest) (*catalog.Restaurant, error)
}

type ItemServiceClientInterface interface {
	GetItemById(request *catalog.GetItemByIdRequest) (*catalog.MenuItem, error)
}
