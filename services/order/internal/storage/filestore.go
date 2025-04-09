package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/yourproject/services/order/internal/model"
)

type FileStore struct {
	mu       sync.Mutex
	filePath string
}

func NewFileStore(dataDir string) *FileStore {
	return &FileStore{
		filePath: dataDir + "/orders.json",
	}
}

func (fs *FileStore) SaveOrder(order *model.Order) (*model.Order, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	orders, err := fs.readOrders()
	if err != nil {
		return nil, err
	}

	orders = append(orders, *order)
	if err := fs.writeOrders(orders); err != nil {
		return nil, err
	}

	return order, nil
}

func (fs *FileStore) GetOrder(id string) (*model.Order, error) {
	orders, err := fs.readOrders()
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		if order.ID == id {
			return &order, nil
		}
	}

	return nil, ErrOrderNotFound
}

// Similar implementations for UpdateOrderStatus and other methods
