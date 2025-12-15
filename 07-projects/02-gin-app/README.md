# Gin Web API

这是一个使用 Gin 框架构建的现代化微服务 Web API 示例，采用分层架构设计（Domain Driven Design）。

## 项目特点

- ✅ 完整的分层架构（Handler → Service → Repository）
- ✅ RESTful API 设计
- ✅ 中间件支持（日志、CORS）
- ✅ 数据验证
- ✅ Web UI 管理界面
- ✅ Docker 支持
- ✅ Makefile 简化操作
- ✅ 可扩展的项目结构

## 项目结构

```
gin-app/
├── cmd/                      # 程序入口
│   └── server/
│       └── main.go          # 主程序
├── internal/                # 内部包（不对外暴露）
│   ├── handler/             # HTTP 处理器
│   │   └── user.go
│   ├── middleware/          # 中间件
│   │   ├── logger.go
│   │   └── cors.go
│   ├── model/               # 数据模型
│   │   └── user.go
│   ├── repository/          # 数据访问层
│   │   └── user.go
│   └── service/             # 业务逻辑层
│       └── user.go
├── pkg/                     # 公共包
│   ├── config/              # 配置管理
│   │   └── config.go
│   └── database/            # 数据库
│       └── database.go
├── api/                     # API 文档
│   └── swagger/
│       └── swagger.md
├── web/                     # 前端文件
│   ├── static/              # 静态资源
│   └── template/            # HTML 模板
│       └── index.html
├── tests/                   # 测试文件
├── Dockerfile               # Docker 配置
├── Makefile                 # 构建文件
├── go.mod                   # Go 模块文件
└── README.md               # 项目说明
```

## 快速开始

### 前置要求

- Go 1.21 或更高版本
- 可选：Docker（用于容器化）
- 可选：Make（用于简化命令）

### 1. 本地运行

```bash
# 进入项目目录
cd gin-app

# 下载依赖
go mod download

# 运行应用
go run cmd/server/main.go
```

应用将在 `http://localhost:8080` 启动

### 2. 使用 Makefile

```bash
# 查看所有可用命令
make help

# 构建应用
make build

# 运行应用
make run

# 开发模式运行（需要安装 air）
make dev

# 清理编译输出
make clean
```

### 3. Docker 运行

```bash
# 构建 Docker 镜像
docker build -t gin-app:latest .

# 运行容器
docker run -p 8080:8080 gin-app:latest
```

## API 使用示例

### 访问 Web UI

打开浏览器访问: `http://localhost:8080`

这是一个完整的用户管理 UI，可以：
- 创建新用户
- 查看用户列表
- 删除用户

### 使用 curl 命令

```bash
# 健康检查
curl http://localhost:8080/health

# 获取用户列表
curl http://localhost:8080/api/v1/users

# 创建用户
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "email": "zhangsan@example.com",
    "phone": "13800138000"
  }'

# 获取特定用户（ID=1）
curl http://localhost:8080/api/v1/users/1

# 更新用户
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三（已更新）",
    "email": "zhangsan.new@example.com"
  }'

# 删除用户
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## 架构说明

### 分层架构

1. **Handler 层** (`internal/handler/`)
   - 处理 HTTP 请求和响应
   - 参数验证和绑定
   - 调用 Service 层

2. **Service 层** (`internal/service/`)
   - 包含业务逻辑
   - 协调 Repository 层
   - 不直接依赖 HTTP

3. **Repository 层** (`internal/repository/`)
   - 数据访问和持久化
   - 数据库操作
   - 可被 Service 层调用

4. **Model 层** (`internal/model/`)
   - 数据结构定义
   - API 请求/响应结构
   - 数据验证规则

5. **Middleware 层** (`internal/middleware/`)
   - 日志中间件
   - CORS 处理
   - 认证授权（可扩展）

## 环境变量

```
PORT=8080              # 服务端口（默认：8080）
ENV=development        # 环境（development/production，默认：development）
DB_DRIVER=sqlite       # 数据库驱动（默认：sqlite）
DATABASE_URL=test.db   # 数据库连接字符串
```

## 扩展建议

### 数据库集成
目前使用内存存储，可以集成真实数据库：
- PostgreSQL（推荐）
- MySQL
- SQLite
- MongoDB

### 认证授权
- JWT Token 认证
- Role Based Access Control (RBAC)
- OAuth2 集成

### 日志和监控
- 结构化日志（zap/logrus）
- Prometheus 指标
- Tracing（Jaeger）

### 测试
```bash
# 运行测试
go test -v ./...

# 生成覆盖率报告
go test -cover ./...
```

### 代码质量
```bash
# 代码格式化
go fmt ./...

# 静态检查（需要安装 golangci-lint）
golangci-lint run
```

## 常见问题

**Q: 如何修改端口？**
A: 设置 `PORT` 环境变量或在 `go run` 时指定：
```bash
PORT=3000 go run cmd/server/main.go
```

**Q: 如何集成数据库？**
A: 修改 `internal/repository/user.go` 中的数据访问逻辑，改为使用真实的 DB 连接（如 GORM）。

**Q: 如何添加新的 API 端点？**
A:
1. 在 `internal/model/` 中定义数据结构
2. 在 `internal/handler/` 中添加处理函数
3. 在 `internal/service/` 中实现业务逻辑
4. 在 `cmd/server/main.go` 中注册路由

## 许可证

MIT License

## 联系方式

如有问题或建议，欢迎提出 Issue。
