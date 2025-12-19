package service

import (
	"errors"

	"ecommerce-platform/internal/models"
	"ecommerce-platform/internal/repository"
	"ecommerce-platform/pkg/utils"
)

// UserService 用户服务接口
type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetUser(id int64) (*models.User, error)
	Login(email, password string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int64) error
	ListUsers() ([]*models.User, error)
}

// userService 用户服务实现
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建新的用户服务
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// 检查邮箱是否已存在
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// 密码加密
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
		Address:  req.Address,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUser(id int64) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id int64) error {
	return s.userRepo.Delete(id)
}

func (s *userService) ListUsers() ([]*models.User, error) {
	return s.userRepo.List()
}
