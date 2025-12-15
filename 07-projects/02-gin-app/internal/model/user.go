package model

import "time"

// User 用户数据模型
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required" gorm:"index"`
	Email     string    `json:"email" binding:"required,email" gorm:"index"`
	Phone     string    `json:"phone"`
	Status    int       `json:"status" gorm:"default:1"` // 1:正常 0:禁用
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email" binding:"email"`
	Phone  string `json:"phone"`
	Status int    `json:"status"`
}

// Response 通用响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
