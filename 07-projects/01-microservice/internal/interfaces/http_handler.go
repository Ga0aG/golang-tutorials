package interfaces

import (
	"encoding/json"
	"microservice/internal/application"
	"net/http"
)

// UserResponse API响应结构
type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Router HTTP路由器
type Router struct {
	userService *application.UserService
	mux         *http.ServeMux
}

// NewRouter 创建新路由器
func NewRouter(userService *application.UserService) *Router {
	r := &Router{
		userService: userService,
		mux:         http.NewServeMux(),
	}
	r.setupRoutes()
	return r
}

// setupRoutes 设置路由
func (r *Router) setupRoutes() {
	r.mux.HandleFunc("/health", r.HealthCheck)
	r.mux.HandleFunc("/api/users", r.ListUsers)
	r.mux.HandleFunc("/api/user", r.HandleUser)
}

// HealthCheck 健康检查
func (r *Router) HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// ListUsers 列出所有用户
func (r *Router) ListUsers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := r.userService.ListUsers()
	responses := make([]UserResponse, len(users))
	for i, u := range users {
		responses[i] = UserResponse{ID: u.ID, Name: u.Name, Email: u.Email}
	}
	json.NewEncoder(w).Encode(responses)
}

// HandleUser 处理单个用户请求
func (r *Router) HandleUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method == http.MethodPost {
		var payload struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		json.NewDecoder(req.Body).Decode(&payload)
		user := r.userService.CreateUser(payload.ID, payload.Name, payload.Email)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(UserResponse{ID: user.ID, Name: user.Name, Email: user.Email})
	} else if req.Method == http.MethodGet {
		id := req.URL.Query().Get("id")
		user := r.userService.GetUser(id)
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
			return
		}
		json.NewEncoder(w).Encode(UserResponse{ID: user.ID, Name: user.Name, Email: user.Email})
	}
}

// Start 启动HTTP服务器
func (r *Router) Start(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}
