package models

import "time"

// OrderStatus 订单状态
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"   // 待支付
	OrderStatusPaid      OrderStatus = "paid"      // 已支付
	OrderStatusShipped   OrderStatus = "shipped"   // 已发货
	OrderStatusDelivered OrderStatus = "delivered" // 已送达
	OrderStatusCancelled OrderStatus = "cancelled" // 已取消
)

// Order 订单模型
type Order struct {
	ID         int64       `json:"id"`
	UserID     int64       `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Status     OrderStatus `json:"status"`
	Items      []OrderItem `json:"items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

// OrderItem 订单项
type OrderItem struct {
	ID        int64   `json:"id"`
	OrderID   int64   `json:"order_id"`
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // 下单时的价格
	Product   Product `json:"product,omitempty"`
}

// CreateOrderRequest 创建订单请求
type CreateOrderRequest struct {
	UserID int64           `json:"user_id" binding:"required"`
	Items  []OrderItemReq  `json:"items" binding:"required,min=1,dive"`
}

// OrderItemReq 订单项请求
type OrderItemReq struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int   `json:"quantity" binding:"required,gt=0"`
}
