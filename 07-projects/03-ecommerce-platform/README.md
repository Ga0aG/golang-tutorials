# E-commerce Platform

一个使用 Go 语言实现的简易电商平台，展示了标准的 Go 项目结构和分层架构设计。

## 项目结构

```
ecommerce-platform/
├── cmd/
│   └── api/
│       └── main.go           # 应用程序入口
├── internal/                 # 私有应用代码
│   ├── models/              # 数据模型
│   │   ├── user.go
│   │   ├── product.go
│   │   └── order.go
│   ├── repository/          # 数据访问层
│   │   ├── user_repository.go
│   │   ├── product_repository.go
│   │   └── order_repository.go
│   ├── service/             # 业务逻辑层
│   │   ├── user_service.go
│   │   ├── product_service.go
│   │   └── order_service.go
│   └── handler/             # HTTP 处理器
│       ├── user_handler.go
│       ├── product_handler.go
│       ├── order_handler.go
│       └── response.go
├── pkg/                     # 可重用的公共代码
│   └── utils/
│       └── password.go      # 密码加密工具
├── config/                  # 配置文件
│   └── config.go
├── go.mod
└── README.md
```

## 架构说明

### 分层架构

1. **Models（模型层）**：定义数据结构和领域模型
2. **Repository（仓储层）**：数据访问和持久化
3. **Service（服务层）**：业务逻辑处理
4. **Handler（处理器层）**：HTTP 请求处理

### 设计模式

- **依赖注入**：各层通过接口依赖，便于测试和扩展
- **仓储模式**：抽象数据访问，当前使用内存存储，可轻松切换到数据库
- **MVC 架构**：清晰的职责分离

## 功能特性

### 用户模块
- 用户注册（密码加密）
- 用户登录
- 用户信息查询

### 商品模块
- 商品创建
- 商品查询（单个/列表/分类）
- 商品更新
- 商品删除
- 库存管理

### 订单模块
- 创建订单（自动库存扣减）
- 订单查询（单个/用户订单/全部）
- 订单状态更新
- 取消订单（自动库存恢复）

## 快速开始

### 安装依赖

```bash
cd ecommerce-platform
go mod download
```

### 运行项目

```bash
go run cmd/api/main.go
```

服务器将在 `http://localhost:8080` 启动。

### 环境变量配置

可以通过环境变量配置服务器：

```bash
export SERVER_PORT=8080
export DB_HOST=localhost
export DB_PORT=5432
```

## API 接口文档

### 用户接口

#### 创建用户
```bash
POST /api/users
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123",
  "phone": "1234567890",
  "address": "123 Main St"
}
```

#### 用户登录
```bash
POST /api/users/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

#### 获取用户
```bash
GET /api/users?id=1
```

#### 列出所有用户
```bash
GET /api/users
```

### 商品接口

#### 创建商品
```bash
POST /api/products
Content-Type: application/json

{
  "name": "iPhone 15",
  "description": "Latest iPhone model",
  "price": 999.99,
  "stock": 100,
  "category": "electronics",
  "image_url": "https://example.com/iphone.jpg"
}
```

#### 获取商品
```bash
GET /api/products?id=1
```

#### 列出所有商品
```bash
GET /api/products
```

#### 按分类查询商品
```bash
GET /api/products?category=electronics
```

#### 更新商品
```bash
PUT /api/products?id=1
Content-Type: application/json

{
  "name": "iPhone 15 Pro",
  "price": 1199.99,
  "stock": 50
}
```

#### 删除商品
```bash
DELETE /api/products?id=1
```

### 订单接口

#### 创建订单
```bash
POST /api/orders
Content-Type: application/json

{
  "user_id": 1,
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 2,
      "quantity": 1
    }
  ]
}
```

#### 获取订单
```bash
GET /api/orders?id=1
```

#### 获取用户订单
```bash
GET /api/orders?user_id=1
```

#### 列出所有订单
```bash
GET /api/orders
```

#### 更新订单状态
```bash
PUT /api/orders/status?id=1
Content-Type: application/json

{
  "status": "paid"
}
```

订单状态：`pending`（待支付）、`paid`（已支付）、`shipped`（已发货）、`delivered`（已送达）、`cancelled`（已取消）

#### 取消订单
```bash
POST /api/orders/cancel?id=1
```

## 测试示例

### 完整流程测试

```bash
# 1. 创建用户
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "email": "alice@example.com",
    "password": "password123",
    "phone": "1234567890",
    "address": "123 Main St"
  }'

# 2. 创建商品
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "price": 1299.99,
    "stock": 50,
    "category": "electronics",
    "image_url": "https://example.com/laptop.jpg"
  }'

# 3. 查看商品列表
curl http://localhost:8080/api/products

# 4. 创建订单
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "items": [
      {
        "product_id": 1,
        "quantity": 1
      }
    ]
  }'

# 5. 查看订单
curl http://localhost:8080/api/orders?id=1

# 6. 更新订单状态
curl -X PUT http://localhost:8080/api/orders/status?id=1 \
  -H "Content-Type: application/json" \
  -d '{"status": "paid"}'
```

## 扩展建议

当前版本使用内存存储数据，适合学习和演示。生产环境建议：

1. **数据库集成**：替换内存仓储为 PostgreSQL/MySQL
2. **身份认证**：添加 JWT 或 Session 认证
3. **日志系统**：集成结构化日志
4. **错误处理**：完善错误处理和验证
5. **中间件**：添加 CORS、限流、日志中间件
6. **单元测试**：为各层编写测试
7. **API 文档**：集成 Swagger/OpenAPI
8. **容器化**：添加 Dockerfile 和 docker-compose

## 技术栈

- Go 1.21+
- net/http（标准库 HTTP 服务器）
- golang.org/x/crypto（密码加密）

## 许可证

MIT License
