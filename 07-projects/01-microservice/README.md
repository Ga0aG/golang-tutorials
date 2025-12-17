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

# 领域驱动设计 DDD(Domain-Driven Design)

## 架构层次（从外到内）：
┌─────────────────────────────────────────┐
│   接口层 (Interface Layer)               │ ← HTTP/GRPC 控制器
│   - 接收请求，返回响应                     │
│   - 处理序列化/反序列化                    │
├─────────────────────────────────────────┤
│   应用层 (Application Layer)             │ ← 应用服务
│   - 协调领域对象完成用例                   │
│   - 事务管理，外部服务调用                  │
├─────────────────────────────────────────┤
│   领域层 (Domain Layer) ←─ 核心！         │ ← domain/ 目录
│   - 业务实体、值对象、领域服务              │
│   - 业务规则、不变式、领域事件              │
├─────────────────────────────────────────┤
│   基础设施层 (Infrastructure Layer)       │ ← 数据库、消息队列等
│   - 数据持久化                            │
│   - 外部API调用                           │
│   - 消息发送                              │
└─────────────────────────────────────────┘

代码即文档；


## 三个核心特征

1. 贫血模型 vs 充血模型

   ```go
   // ❌ 贫血模型（Anemic Model）- 只有数据，没有行为
   type Order struct {
      ID     int
      UserID int
      Total  float64
      Status string
      // 没有业务方法
      // 业务逻辑在Service层
   }

   // ✅ 充血模型（Rich Model）- 数据 + 行为
   type Order struct {
      id          OrderID
      customer    *Customer
      items       []*OrderItem
      status      OrderStatus

      // 业务行为封装在实体内部
      Place() error
      Cancel(reason string) error
      AddItem(product *Product, quantity int) error
      CalculateTotal() Money
   }
   ```

2. 值对象（Value Object）

   ```go
   // 没有唯一标识，通过属性值判断相等
   type EmailAddress struct {
      value string
   }

   func NewEmailAddress(email string) (*EmailAddress, error) {
      if !isValidEmail(email) {
         return nil, errors.New("无效的邮箱地址")
      }
      return &EmailAddress{value: email}, nil
   }

   func (e EmailAddress) Value() string {
      return e.value
   }

   // 使用
   customerEmail, _ := NewEmailAddress("user@example.com")
   // 这个邮箱地址就是一个值，没有ID，相同邮箱就是同一个对象
   ```

3. 实体（Entity）

   ```go
   // 有唯一标识，即使属性变化还是同一个实体
   type Customer struct {
      id      CustomerID  // 唯一标识
      name    string      // 可以改名字
      email   EmailAddress // 可以改邮箱
      status  CustomerStatus

      // 但ID不变，就还是同一个客户
   }

   // 实体标识
   type CustomerID struct {
      value string  // 可以是UUID或业务ID
   }

   func (id CustomerID) String() string {
      return id.value
   }

   func (id CustomerID) Equals(other CustomerID) bool {
      return id.value == other.value
   }
   ```

##　vs 传统数据库驱动

| 方面         | 数据库驱动       | 领域模型               |
| ------------ | ---------------- | ---------------------- |
| 业务逻辑位置 | Service 层       | 实体和聚合根内部       |
| 代码组织     | 按技术分层       | 按业务概念组织         |
| 可测试性     | 需要数据库       | 可单元测试（mock仓储） |
| 业务复杂性   | 简单 CRUD 适合   | 复杂业务规则适合       |
| 数据一致性   | 数据库事务       | 聚合根边界保证         |
| 演化能力     | 数据库变更影响大 | 业务逻辑独立演化       |

1. 传统crud开发

   ```go
   // controller/order_controller.go
   func CreateOrder(c *gin.Context) {
      // 1. 解析请求
      var req CreateOrderRequest
      c.BindJSON(&req)

      // 2. 直接操作数据库
      tx := db.Begin()

      // 检查库存
      for _, item := range req.Items {
         var stock int
         tx.Raw("SELECT stock FROM products WHERE id = ?", item.ProductID).Scan(&stock)
         if stock < item.Quantity {
               tx.Rollback()
               c.JSON(400, gin.H{"error": "库存不足"})
               return
         }
      }

      // 创建订单
      orderID := generateID()
      tx.Exec("INSERT INTO orders (id, user_id, total, status) VALUES (?, ?, ?, ?)",
         orderID, req.UserID, calculateTotal(req.Items), "pending")

      // 创建订单项
      for _, item := range req.Items {
         tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)",
               orderID, item.ProductID, item.Quantity, getPrice(item.ProductID))
      }

      // 更新库存
      for _, item := range req.Items {
         tx.Exec("UPDATE products SET stock = stock - ? WHERE id = ?",
               item.Quantity, item.ProductID)
      }

      tx.Commit()

      // 3. 返回响应
      c.JSON(200, gin.H{"order_id": orderID})
   }
   ```

2. 领域驱动开发

   ```go
   // handler/order_handler.go
   func (h *OrderHandler) CreateOrder(c *gin.Context) {
      // 1. 解析请求
      var req CreateOrderRequest
      c.BindJSON(&req)

      // 2. 调用应用服务（协调领域对象）
      order, err := h.orderSvc.CreateOrder(
         CustomerID(req.CustomerID),
         req.Items,
      )

      // 3. 返回响应
      if err != nil {
         c.JSON(400, gin.H{"error": err.Error()})
         return
      }

      c.JSON(200, gin.H{
         "order_id": order.ID().String(),
         "status":   order.Status().String(),
         "total":    order.Total().String(),
      })
   }
   ```

   ```go
   // application/order_service.go
   func (s *OrderService) CreateOrder(
      customerID CustomerID,
      items []OrderItemRequest,
   ) (*domain.Order, error) {
      // 使用领域对象执行业务逻辑
      return s.orderDomainService.CreateOrder(customerID, items)
   }
   ```

## 聚合

聚合就是一组相关对象的集合，这些对象必须作为一个整体来维护业务规则的一致性。

三个关键特性:

1. 有一个聚合根（Aggregate Root）- 唯一的入口点
2. 一致性边界 - 内部对象一起保持一致性
3. 事务边界 - 通常在一个事务内保存

没有聚合的问题

```go
// 分散的实体，没有聚合保护
type Order struct {
    ID     int
    Items  []OrderItem  // 可以直接修改
}

type OrderItem struct {
    OrderID int
    ProductID int
    Quantity int
}

// 问题：可以从外部随意修改OrderItem
func main() {
    order := GetOrder(1)

    // ❌ 危险：直接修改内部对象
    order.Items[0].Quantity = 999  // 绕过了业务规则检查！
    // 订单总额可能不一致了！
    // 库存检查可能被跳过！
}
```

✅ 使用聚合保护

```go
// Order 作为聚合根
type Order struct {
    id    OrderID
    items []*OrderItem  // 私有字段
    total Money
}

// 外部只能通过聚合根的方法访问
func (o *Order) AddItem(product *Product, quantity int) error {
    // ✅ 业务规则检查
    if o.status != OrderStatusDraft {
        return errors.New("只能向草稿订单添加商品")
    }

    if quantity <= 0 {
        return errors.New("数量必须大于0")
    }

    // 创建订单项（内部对象）
    item := &OrderItem{
        product:  product,
        quantity: quantity,
        price:    product.CurrentPrice(),
    }

    o.items = append(o.items, item)
    o.total = o.total.Add(item.Subtotal())  // ✅ 保持总额同步

    return nil
}

func (o *Order) RemoveItem(productID ProductID) error {
    // 类似地，有完整的业务逻辑
    // ...
}

// ❌ 外部无法直接访问items
// order.items = ...  // 编译错误：items是私有字段
```

怎么这么多黑话

