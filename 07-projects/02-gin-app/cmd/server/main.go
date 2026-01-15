package main

import (
	"gin-app/internal/handler"
	"gin-app/internal/middleware"
	"gin-app/internal/repository"
	"gin-app/internal/service"
	"gin-app/pkg/config"
	"gin-app/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()
	log.Printf("配置加载完成: Port=%d, Env=%s\n", cfg.Port, cfg.Env)

	// 初始化数据库
	db := database.Init(cfg)
	log.Println("数据库初始化完成")

	// 设置gin模式
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 创建路由
	router := gin.Default()

	// 注册中间件
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// 初始化各层
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 注册API路由
	registerRoutes(router, userHandler)

	// 启动服务
	addr := ":" + cfg.Port
	log.Printf("服务启动在 http://localhost%s\n", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

func registerRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	// 提供静态文件
	router.LoadHTMLGlob("web/template/*.html")
	router.Static("/assets", "./web/assets")

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"code":   0,
		})
	})

	// 根路径 - 返回前端页面
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	// API 根路径信息
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "欢迎使用 Gin Web API",
			"version": "1.0.0",
		})
	})

	// 用户相关API
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", userHandler.ListUsers)
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}
