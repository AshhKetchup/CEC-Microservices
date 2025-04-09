package model

import (
	"time"

	common "cec-project/delivery-app/services/common/gen"
)

type OrderStatus int

const (
	OrderStatusPending OrderStatus = iota
	OrderStatusProcessing
	OrderStatusCompleted
	OrderStatusCancelled
)

type Order struct {
	ID        string
	UserID    string
	Items     []Product
	Total     float64
	Status    OrderStatus
	CreatedAt time.Time
}

type OrderRequest struct {
	UserID     string
	ProductIDs []string
}

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

// Storage interface for dependency inversion
type OrderStorage interface {
	SaveOrder(*Order) (*Order, error)
	GetOrder(string) (*Order, error)
	UpdateOrderStatus(string, OrderStatus) (*Order, error)
}

// Proto conversion helpers
func ConvertToPBOrderStatus(s OrderStatus) common.OrderStatus {
	return common.OrderStatus(s)
}

func ConvertFromPBProducts(pbProducts []*common.ProductResponse) []Product {
	products := make([]Product, len(pbProducts))
	for i, p := range pbProducts {
		products[i] = Product{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		}
	}
	return products
}
