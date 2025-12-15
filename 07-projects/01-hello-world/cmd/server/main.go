package main

import (
	"fmt"
	"log"
	"microservice/internal/application"
	"microservice/internal/interfaces"
	"microservice/pkg/config"
	"microservice/pkg/logger"
)

func main() {
	// 初始化配置
	cfg := config.LoadConfig()
	log.Printf("配置加载完成: %+v\n", cfg)

	// 初始化日志
	logger.Init(cfg.LogLevel)
	logger.Info("微服务启动中...")

	// 初始化应用服务
	appService := application.NewUserService()

	// 初始化路由和接口
	router := interfaces.NewRouter(appService)

	// 启动服务
	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info(fmt.Sprintf("服务启动在 %s", addr))

	if err := router.Start(addr); err != nil {
		logger.Error(fmt.Sprintf("服务启动失败: %v", err))
		panic(err)
	}
}
