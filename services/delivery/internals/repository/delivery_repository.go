package repository

import (
	"context"

	"cec-project/delivery-app/services/delivery/internal/client"
	"cec-project/delivery-app/services/delivery/internal/model"
)

type DeliveryRepository interface {
	CreateDelivery(ctx context.Context, req model.DeliveryRequest) (*model.Delivery, error)
	GetDelivery(ctx context.Context, id string) (*model.Delivery, error)
	UpdateDeliveryStatus(ctx context.Context, id string, status model.DeliveryStatus) (*model.Delivery, error)
}

type DeliveryRepo struct {
	store       model.DeliveryStorage
	orderClient client.OrderClient
}

func NewDeliveryRepository(store model.DeliveryStorage, oc client.OrderClient) *DeliveryRepo {
	return &DeliveryRepo{
		store:       store,
		orderClient: oc,
	}
}

func (r *DeliveryRepo) CreateDelivery(ctx context.Context, req model.DeliveryRequest) (*model.Delivery, error) {
	// Validate order exists
	_, err := r.orderClient.GetOrder(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}

	delivery := &model.Delivery{
		OrderID:           req.OrderID,
		Address:           req.Address,
		Status:            model.DeliveryStatusPending,
		EstimatedDelivery: time.Now().Add(24 * time.Hour),
	}

	return r.store.SaveDelivery(delivery)
}
