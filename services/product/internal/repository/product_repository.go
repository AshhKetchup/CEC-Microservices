package repository

import (
	"cec-project/delivery-app/services/product/internal/model"
	"cec-project/delivery-app/services/product/internal/storage"
	"context"
	"github.com/google/uuid"
	"time"
)

type ProductRepository struct {
	store *storage.FileStore
}

func NewProductRepository(store *storage.FileStore) ProductRepository {
	return &ProductRepository{store: store}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product model.Product) (model.Product, error) {
	product.ID = uuid.New().String()
	product.CreatedAt = time.Now()

	err := r.store.Save(product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (r *productRepository) GetProduct(ctx context.Context, id string) (model.Product, error) {
	return r.store.Load(id)
}
