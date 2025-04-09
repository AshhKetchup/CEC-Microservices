package model

import (
	"time"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
}

type ProductStorage interface {
	SaveProduct(*Product) error
	GetProduct(string) (*Product, error)
	ListProducts(int, int) ([]Product, error)
	UpdateProduct(*Product) error
	DeleteProduct(string) error
}
