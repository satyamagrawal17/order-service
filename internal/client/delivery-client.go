package client

import "ordering_service/pb/delivery"

type DeliveryServiceClientInterface interface {
	AssignDelivery(*delivery.AssignDeliveryRequest) (*delivery.AssignDeliveryResponse, error)
	UpdateDeliveryStatus(*delivery.UpdateDeliveryStatusRequest) (*delivery.UpdateDeliveryStatusResponse, error)
}
