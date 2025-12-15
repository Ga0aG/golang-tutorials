package service

import (
	"gin-app/internal/model"
	"gin-app/internal/repository"
	"time"
)

// UserService 用户业务逻辑层
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService 创建用户服务
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	user := &model.User{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser 获取用户
func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.repo.GetByID(id)
}

// ListUsers 获取用户列表
func (s *UserService) ListUsers() ([]model.User, error) {
	return s.repo.List()
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(id int, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	updateData := &model.User{
		Name:   req.Name,
		Email:  req.Email,
		Phone:  req.Phone,
		Status: req.Status,
	}

	err = s.repo.Update(id, updateData)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
