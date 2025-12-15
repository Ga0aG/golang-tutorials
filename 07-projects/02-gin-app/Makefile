.PHONY: help build run clean test docker-build docker-run

# 变量定义
APP_NAME = gin-app
PORT = 8080
DOCKER_IMAGE = $(APP_NAME):latest

help:
	@echo "Gin Web API 项目"
	@echo "================="
	@echo "可用命令:"
	@echo "  make build          构建应用"
	@echo "  make run            运行应用"
	@echo "  make dev            开发模式运行（自动重载）"
	@echo "  make clean          清理编译输出"
	@echo "  make test           运行测试"
	@echo "  make docker-build   构建Docker镜像"
	@echo "  make docker-run     运行Docker容器"
	@echo "  make fmt            代码格式化"
	@echo "  make lint           代码静态检查"

build:
	@echo "编译应用..."
	go build -o bin/$(APP_NAME) cmd/server/main.go
	@echo "编译完成: bin/$(APP_NAME)"

run: build
	@echo "运行应用..."
	PORT=$(PORT) ./bin/$(APP_NAME)

dev:
	@echo "开发模式运行..."
	@which air > /dev/null || (echo "安装air..." && go install github.com/cosmtrek/air@latest)
	air

clean:
	@echo "清理编译输出..."
	rm -rf bin/
	@echo "清理完成"

test:
	@echo "运行测试..."
	go test -v ./...

fmt:
	@echo "代码格式化..."
	go fmt ./...

lint:
	@echo "代码静态检查..."
	@which golangci-lint > /dev/null || (echo "请安装 golangci-lint"; exit 1)
	golangci-lint run

docker-build:
	@echo "构建Docker镜像..."
	docker build -t $(DOCKER_IMAGE) .
	@echo "镜像构建完成: $(DOCKER_IMAGE)"

docker-run: docker-build
	@echo "运行Docker容器..."
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE)

docker-stop:
	@echo "停止Docker容器..."
	docker ps | grep $(APP_NAME) | awk '{print $$1}' | xargs docker stop

install-deps:
	@echo "下载依赖..."
	go mod download
	go mod tidy
	@echo "依赖安装完成"

.DEFAULT_GOAL := help
