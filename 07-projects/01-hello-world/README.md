# 微服务项目

这是一个标准的微服务项目模板，采用DDD（领域驱动设计）和分层架构。

## 项目结构

```
microservice/
├── cmd/                          # 命令行应用
│   └── server/
│       └── main.go              # 程序入口
├── internal/                     # 内部包（不向外部公开）
│   ├── domain/                  # 领域模型
│   │   └── user.go             # User 聚合根
│   ├── application/             # 应用服务层
│   │   └── user_service.go     # 业务逻辑
│   ├── infrastructure/          # 基础设施层
│   │   └── database.go         # 数据库实现
│   └── interfaces/              # 接口适配层
│       └── http_handler.go     # HTTP路由和处理
├── pkg/                          # 公共包
│   ├── config/                  # 配置管理
│   └── logger/                  # 日志系统
├── deployments/                 # 部署配置
│   ├── docker-compose.yml       # Docker Compose 配置
│   └── k8s-deployment.yml       # Kubernetes 配置
├── go.mod                        # Go 模块文件
├── .gitignore                    # Git 忽略文件
├── Dockerfile                    # Docker 镜像构建
└── README.md                     # 项目文档
```

## 层级说明

### Domain（领域层）
- 包含业务逻辑的实体和值对象
- 不依赖任何框架或技术细节
- 例：`User` 实体

### Application（应用层）
- 协调域对象完成业务流程
- 不包含业务规则，只协调
- 例：`UserService` 应用服务

### Infrastructure（基础设施层）
- 实现外部系统集成
- 数据库、消息队列、第三方API等
- 例：`Database` 数据库连接

### Interfaces（接口层）
- 暴露给外部的API
- HTTP、gRPC、消息队列等
- 例：`HTTPHandler` HTTP处理器

## 运行项目

### 本地运行

```bash
# 进入项目目录
cd microservice

# 运行应用
go run cmd/server/main.go
```

服务将在 `http://localhost:8080` 启动

### API 端点

```bash
# 健康检查
curl http://localhost:8080/health

# 列出所有用户
curl http://localhost:8080/api/users

# 创建用户
curl -X POST http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  -d '{"id":"1","name":"John","email":"john@example.com"}'

# 获取用户
curl http://localhost:8080/api/user?id=1
```

### Docker 运行

```bash
# 构建镜像
docker build -t microservice .

# 运行容器
docker run -p 8080:8080 microservice
```

### Docker Compose 运行

```bash
cd deployments
docker-compose up
```

### Kubernetes 部署

```bash
# 应用配置
kubectl apply -f deployments/k8s-deployment.yml

# 查看部署状态
kubectl get deployments
kubectl get pods
kubectl get svc
```

## 环境变量

```
PORT=8080              # 服务端口
LOG_LEVEL=info         # 日志级别
DATABASE_URL=...       # 数据库连接字符串
ENV=development        # 环境（development/production）
```

## 开发约定

1. **Package 组织**
   - `cmd/` - 可执行程序
   - `internal/` - 内部包，不允许外部导入
   - `pkg/` - 公共包，可被其他项目使用

2. **命名规范**
   - 包名使用小写
   - 导出的标识符使用PascalCase
   - 文件名使用snake_case

3. **依赖注入**
   - 通过构造函数传递依赖
   - 便于测试和解耦

## 扩展建议

- 添加数据库连接（PostgreSQL/MySQL）
- 集成 gRPC 支持
- 添加单元测试和集成测试
- 实现中间件（认证、日志、追踪）
- 添加配置文件支持（YAML/TOML）
- 集成 Prometheus 监控指标
