package service

import (
	"errors"

	"ecommerce-platform/internal/models"
	"ecommerce-platform/internal/repository"
)

// OrderService 订单服务接口
type OrderService interface {
	CreateOrder(req *models.CreateOrderRequest) (*models.Order, error)
	GetOrder(id int64) (*models.Order, error)
	GetUserOrders(userID int64) ([]*models.Order, error)
	UpdateOrderStatus(id int64, status models.OrderStatus) error
	CancelOrder(id int64) error
	ListOrders() ([]*models.Order, error)
}

// orderService 订单服务实现
type orderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
	userRepo    repository.UserRepository
}

// NewOrderService 创建新的订单服务
func NewOrderService(
	orderRepo repository.OrderRepository,
	productRepo repository.ProductRepository,
	userRepo repository.UserRepository,
) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (s *orderService) CreateOrder(req *models.CreateOrderRequest) (*models.Order, error) {
	// 验证用户是否存在
	_, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	var orderItems []models.OrderItem
	var totalPrice float64

	// 处理订单项
	for _, item := range req.Items {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found: " + err.Error())
		}

		// 检查库存
		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		// 减少库存
		if err := s.productRepo.UpdateStock(product.ID, -item.Quantity); err != nil {
			return nil, err
		}

		orderItem := models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
			Product:   *product,
		}
		orderItems = append(orderItems, orderItem)
		totalPrice += product.Price * float64(item.Quantity)
	}

	order := &models.Order{
		UserID:     req.UserID,
		TotalPrice: totalPrice,
		Status:     models.OrderStatusPending,
		Items:      orderItems,
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) GetOrder(id int64) (*models.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *orderService) GetUserOrders(userID int64) ([]*models.Order, error) {
	return s.orderRepo.GetByUserID(userID)
}

func (s *orderService) UpdateOrderStatus(id int64, status models.OrderStatus) error {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	order.Status = status
	return s.orderRepo.Update(order)
}

func (s *orderService) CancelOrder(id int64) error {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	if order.Status != models.OrderStatusPending {
		return errors.New("only pending orders can be cancelled")
	}

	// 恢复库存
	for _, item := range order.Items {
		if err := s.productRepo.UpdateStock(item.ProductID, item.Quantity); err != nil {
			return err
		}
	}

	order.Status = models.OrderStatusCancelled
	return s.orderRepo.Update(order)
}

func (s *orderService) ListOrders() ([]*models.Order, error) {
	return s.orderRepo.List()
}
