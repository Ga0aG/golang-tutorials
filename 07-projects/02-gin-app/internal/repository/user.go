package repository

import (
	"gin-app/internal/model"
)

// 模拟数据库存储
var users = make(map[int]*model.User)
var nextID = 1

// UserRepository 用户数据访问层
type UserRepository struct {
	db interface{} // 可以替换为真实数据库连接
}

// NewUserRepository 创建用户仓储
func NewUserRepository(db interface{}) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create 创建用户
func (r *UserRepository) Create(user *model.User) error {
	user.ID = nextID
	users[nextID] = user
	nextID++
	return nil
}

// GetByID 根据ID获取用户
func (r *UserRepository) GetByID(id int) (*model.User, error) {
	if user, ok := users[id]; ok {
		return user, nil
	}
	return nil, nil
}

// List 获取所有用户
func (r *UserRepository) List() ([]model.User, error) {
	result := make([]model.User, 0)
	for _, user := range users {
		result = append(result, *user)
	}
	return result, nil
}

// Update 更新用户
func (r *UserRepository) Update(id int, user *model.User) error {
	if existing, ok := users[id]; ok {
		if user.Name != "" {
			existing.Name = user.Name
		}
		if user.Email != "" {
			existing.Email = user.Email
		}
		if user.Phone != "" {
			existing.Phone = user.Phone
		}
		if user.Status != 0 {
			existing.Status = user.Status
		}
		return nil
	}
	return nil
}

// Delete 删除用户
func (r *UserRepository) Delete(id int) error {
	delete(users, id)
	return nil
}
