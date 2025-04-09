package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"cec-project/delivery-app/services/delivery/internal/model"
)

type FileStore struct {
	mu       sync.Mutex
	filePath string
}

func NewFileStore(dataDir string) *FileStore {
	return &FileStore{
		filePath: dataDir + "/deliveries.json",
	}
}

func (fs *FileStore) SaveDelivery(delivery *model.Delivery) (*model.Delivery, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	deliveries, err := fs.readDeliveries()
	if err != nil {
		return nil, err
	}

	delivery.ID = generateID()
	delivery.CreatedAt = time.Now()
	deliveries = append(deliveries, *delivery)

	if err := fs.writeDeliveries(deliveries); err != nil {
		return nil, err
	}

	return delivery, nil
}

func (fs *FileStore) GetDelivery(id string) (*model.Delivery, error) {
	deliveries, err := fs.readDeliveries()
	if err != nil {
		return nil, err
	}

	for _, d := range deliveries {
		if d.ID == id {
			return &d, nil
		}
	}
	return nil, model.ErrDeliveryNotFound
}

func (fs *FileStore) readDeliveries() ([]model.Delivery, error) {
	if _, err := os.Stat(fs.filePath); os.IsNotExist(err) {
		return []model.Delivery{}, nil
	}

	file, err := os.Open(fs.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var deliveries []model.Delivery
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&deliveries)
	return deliveries, err
}

func (fs *FileStore) writeDeliveries(deliveries []model.Delivery) error {
	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(deliveries)
}
