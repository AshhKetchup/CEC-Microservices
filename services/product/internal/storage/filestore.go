package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"cec-project/delivery-app/services/product/internal/model"
)

type FileStore struct {
	mu       sync.Mutex
	filePath string
}

func NewFileStore(dataDir string) *FileStore {
	return &FileStore{
		filePath: dataDir + "/products.json",
	}
}

func (fs *FileStore) SaveProduct(product *model.Product) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	products, err := fs.readProducts()
	if err != nil {
		return err
	}

	product.ID = generateID()
	product.CreatedAt = time.Now()
	products = append(products, *product)

	return fs.writeProducts(products)
}

func (fs *FileStore) GetProduct(id string) (*model.Product, error) {
	products, err := fs.readProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, ErrProductNotFound
}

func (fs *FileStore) readProducts() ([]model.Product, error) {
	var products []model.Product

	if _, err := os.Stat(fs.filePath); os.IsNotExist(err) {
		return []model.Product{}, nil
	}

	file, err := os.Open(fs.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	return products, err
}

func (fs *FileStore) writeProducts(products []model.Product) error {
	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(products)
}
