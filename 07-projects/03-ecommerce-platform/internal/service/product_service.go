package service

import (
	"ecommerce-platform/internal/models"
	"ecommerce-platform/internal/repository"
)

// ProductService 商品服务接口
type ProductService interface {
	CreateProduct(req *models.CreateProductRequest) (*models.Product, error)
	GetProduct(id int64) (*models.Product, error)
	UpdateProduct(id int64, req *models.UpdateProductRequest) (*models.Product, error)
	DeleteProduct(id int64) error
	ListProducts() ([]*models.Product, error)
	GetProductsByCategory(category string) ([]*models.Product, error)
}

// productService 商品服务实现
type productService struct {
	productRepo repository.ProductRepository
}

// NewProductService 创建新的商品服务
func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(req *models.CreateProductRequest) (*models.Product, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) GetProduct(id int64) (*models.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *productService) UpdateProduct(id int64, req *models.UpdateProductRequest) (*models.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}
	if req.Category != "" {
		product.Category = req.Category
	}
	if req.ImageURL != "" {
		product.ImageURL = req.ImageURL
	}

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) DeleteProduct(id int64) error {
	return s.productRepo.Delete(id)
}

func (s *productService) ListProducts() ([]*models.Product, error) {
	return s.productRepo.List()
}

func (s *productService) GetProductsByCategory(category string) ([]*models.Product, error) {
	return s.productRepo.GetByCategory(category)
}
