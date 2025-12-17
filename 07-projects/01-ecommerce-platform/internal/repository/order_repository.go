package repository

import (
	"errors"
	"sync"
	"time"

	"ecommerce-platform/internal/models"
)

// OrderRepository 订单仓储接口
type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id int64) (*models.Order, error)
	GetByUserID(userID int64) ([]*models.Order, error)
	Update(order *models.Order) error
	Delete(id int64) error
	List() ([]*models.Order, error)
}

// InMemoryOrderRepository 内存实现的订单仓储
type InMemoryOrderRepository struct {
	orders map[int64]*models.Order
	nextID int64
	mu     sync.RWMutex
}

// NewInMemoryOrderRepository 创建新的订单仓储
func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[int64]*models.Order),
		nextID: 1,
	}
}

func (r *InMemoryOrderRepository) Create(order *models.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	order.ID = r.nextID
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	r.orders[order.ID] = order
	r.nextID++
	return nil
}

func (r *InMemoryOrderRepository) GetByID(id int64) (*models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, exists := r.orders[id]
	if !exists {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (r *InMemoryOrderRepository) GetByUserID(userID int64) ([]*models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	orders := make([]*models.Order, 0)
	for _, order := range r.orders {
		if order.UserID == userID {
			orders = append(orders, order)
		}
	}
	return orders, nil
}

func (r *InMemoryOrderRepository) Update(order *models.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[order.ID]; !exists {
		return errors.New("order not found")
	}
	order.UpdatedAt = time.Now()
	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[id]; !exists {
		return errors.New("order not found")
	}
	delete(r.orders, id)
	return nil
}

func (r *InMemoryOrderRepository) List() ([]*models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	orders := make([]*models.Order, 0, len(r.orders))
	for _, order := range r.orders {
		orders = append(orders, order)
	}
	return orders, nil
}
