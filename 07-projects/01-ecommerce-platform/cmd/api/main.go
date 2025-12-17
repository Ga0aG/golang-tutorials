package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce-platform/config"
	"ecommerce-platform/internal/handler"
	"ecommerce-platform/internal/repository"
	"ecommerce-platform/internal/service"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化仓储层
	userRepo := repository.NewInMemoryUserRepository()
	productRepo := repository.NewInMemoryProductRepository()
	orderRepo := repository.NewInMemoryOrderRepository()

	// 初始化服务层
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	orderService := service.NewOrderService(orderRepo, productRepo, userRepo)

	// 初始化处理器
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)
	orderHandler := handler.NewOrderHandler(orderService)

	// 设置路由
	mux := http.NewServeMux()

	// 用户路由
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		case http.MethodGet:
			if r.URL.Query().Get("id") != "" {
				userHandler.GetUser(w, r)
			} else {
				userHandler.ListUsers(w, r)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/users/login", userHandler.Login)

	// 商品路由
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			productHandler.CreateProduct(w, r)
		case http.MethodGet:
			if r.URL.Query().Get("id") != "" {
				productHandler.GetProduct(w, r)
			} else {
				productHandler.ListProducts(w, r)
			}
		case http.MethodPut:
			productHandler.UpdateProduct(w, r)
		case http.MethodDelete:
			productHandler.DeleteProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 订单路由
	mux.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			orderHandler.CreateOrder(w, r)
		case http.MethodGet:
			if r.URL.Query().Get("id") != "" {
				orderHandler.GetOrder(w, r)
			} else if r.URL.Query().Get("user_id") != "" {
				orderHandler.GetUserOrders(w, r)
			} else {
				orderHandler.ListOrders(w, r)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/orders/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			orderHandler.UpdateOrderStatus(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/orders/cancel", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			orderHandler.CancelOrder(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 首页
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to E-commerce Platform API"))
	})

	// 启动服务器
	addr := ":" + cfg.ServerPort
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	fmt.Println("API Endpoints:")
	fmt.Println("  POST   /api/users - Create user")
	fmt.Println("  GET    /api/users?id=1 - Get user by ID")
	fmt.Println("  GET    /api/users - List all users")
	fmt.Println("  POST   /api/users/login - User login")
	fmt.Println("  POST   /api/products - Create product")
	fmt.Println("  GET    /api/products?id=1 - Get product by ID")
	fmt.Println("  GET    /api/products - List all products")
	fmt.Println("  GET    /api/products?category=electronics - List products by category")
	fmt.Println("  PUT    /api/products?id=1 - Update product")
	fmt.Println("  DELETE /api/products?id=1 - Delete product")
	fmt.Println("  POST   /api/orders - Create order")
	fmt.Println("  GET    /api/orders?id=1 - Get order by ID")
	fmt.Println("  GET    /api/orders?user_id=1 - Get user orders")
	fmt.Println("  GET    /api/orders - List all orders")
	fmt.Println("  PUT    /api/orders/status?id=1 - Update order status")
	fmt.Println("  POST   /api/orders/cancel?id=1 - Cancel order")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
