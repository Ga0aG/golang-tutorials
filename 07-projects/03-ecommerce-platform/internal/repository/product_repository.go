package repository

import (
	"errors"
	"sync"
	"time"

	"ecommerce-platform/internal/models"
)

// ProductRepository 商品仓储接口
type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id int64) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int64) error
	List() ([]*models.Product, error)
	GetByCategory(category string) ([]*models.Product, error)
	UpdateStock(id int64, quantity int) error
}

// InMemoryProductRepository 内存实现的商品仓储
type InMemoryProductRepository struct {
	products map[int64]*models.Product
	nextID   int64
	mu       sync.RWMutex
}

// NewInMemoryProductRepository 创建新的商品仓储
func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[int64]*models.Product),
		nextID:   1,
	}
}

func (r *InMemoryProductRepository) Create(product *models.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.ID = r.nextID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	r.products[product.ID] = product
	r.nextID++
	return nil
}

func (r *InMemoryProductRepository) GetByID(id int64) (*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *InMemoryProductRepository) Update(product *models.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[product.ID]; !exists {
		return errors.New("product not found")
	}
	product.UpdatedAt = time.Now()
	r.products[product.ID] = product
	return nil
}

func (r *InMemoryProductRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[id]; !exists {
		return errors.New("product not found")
	}
	delete(r.products, id)
	return nil
}

func (r *InMemoryProductRepository) List() ([]*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]*models.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

func (r *InMemoryProductRepository) GetByCategory(category string) ([]*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]*models.Product, 0)
	for _, product := range r.products {
		if product.Category == category {
			products = append(products, product)
		}
	}
	return products, nil
}

func (r *InMemoryProductRepository) UpdateStock(id int64, quantity int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	product, exists := r.products[id]
	if !exists {
		return errors.New("product not found")
	}

	newStock := product.Stock + quantity
	if newStock < 0 {
		return errors.New("insufficient stock")
	}

	product.Stock = newStock
	product.UpdatedAt = time.Now()
	return nil
}
