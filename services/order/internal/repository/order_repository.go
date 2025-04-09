package repository

import (
	"context"

	"cec-project/delivery-app/services/order/internal/model"
	"cec-project/delivery-appservices/order/internal/client"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req model.OrderRequest) (*model.Order, error)
	GetOrder(ctx context.Context, id string) (*model.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status model.OrderStatus) (*model.Order, error)
}

type OrderRepo struct {
	store         model.OrderStorage
	productClient client.ProductClient
}

func NewOrderRepository(store model.OrderStorage, pc client.ProductClient) *OrderRepo {
	return &OrderRepo{
		store:         store,
		productClient: pc,
	}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, req model.OrderRequest) (*model.Order, error) {
	// Fetch products from product service
	products, total, err := r.productClient.GetProducts(ctx, req.ProductIDs)
	if err != nil {
		return nil, err
	}

	order := &model.Order{
		UserID: req.UserID,
		Items:  products,
		Total:  total,
		Status: model.OrderStatusPending,
	}

	return r.store.SaveOrder(order)
}

func (r *OrderRepo) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	return r.store.GetOrder(id)
}

func (r *OrderRepo) UpdateOrderStatus(ctx context.Context, id string, status model.OrderStatus) (*model.Order, error) {
	return r.store.UpdateOrderStatus(id, status)
}
