package application

import (
	"microservice/internal/domain"
)

// UserService 应用服务 - 处理业务逻辑
type UserService struct {
	users map[string]*domain.User
}

// NewUserService 创建新的用户应用服务
func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*domain.User),
	}
}

// CreateUser 创建用户
func (us *UserService) CreateUser(id, name, email string) *domain.User {
	user := domain.NewUser(id, name, email)
	us.users[id] = user
	return user
}

// GetUser 获取用户
func (us *UserService) GetUser(id string) *domain.User {
	return us.users[id]
}

// ListUsers 列出所有用户
func (us *UserService) ListUsers() []*domain.User {
	users := make([]*domain.User, 0)
	for _, user := range us.users {
		users = append(users, user)
	}
	return users
}
