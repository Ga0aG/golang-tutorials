# Gin Web API 文档

## 基础信息
- **基础URL**: `http://localhost:8080/api/v1`
- **版本**: 1.0.0

## 端点列表

### 健康检查
```
GET /health
```

响应:
```json
{
  "status": "ok",
  "code": 0
}
```

### 获取用户列表
```
GET /api/v1/users
```

响应:
```json
{
  "code": 0,
  "message": "获取成功",
  "data": [
    {
      "id": 1,
      "name": "John",
      "email": "john@example.com",
      "phone": "13800138000",
      "status": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### 创建用户
```
POST /api/v1/users
```

请求体:
```json
{
  "name": "John",
  "email": "john@example.com",
  "phone": "13800138000"
}
```

响应:
```json
{
  "code": 0,
  "message": "创建成功",
  "data": {
    "id": 1,
    "name": "John",
    "email": "john@example.com",
    "phone": "13800138000",
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 获取用户详情
```
GET /api/v1/users/:id
```

参数:
- `id` (path, required): 用户ID

响应:
```json
{
  "code": 0,
  "message": "获取成功",
  "data": {
    "id": 1,
    "name": "John",
    "email": "john@example.com",
    "phone": "13800138000",
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 更新用户
```
PUT /api/v1/users/:id
```

参数:
- `id` (path, required): 用户ID

请求体:
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "13800138001",
  "status": 1
}
```

响应:
```json
{
  "code": 0,
  "message": "更新成功",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "13800138001",
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T01:00:00Z"
  }
}
```

### 删除用户
```
DELETE /api/v1/users/:id
```

参数:
- `id` (path, required): 用户ID

响应:
```
204 No Content
```

## 错误响应

所有错误响应格式:
```json
{
  "code": 400,
  "message": "错误信息描述"
}
```

常见错误代码:
- `400`: 请求参数无效
- `404`: 资源不存在
- `500`: 服务器错误

## 测试示例

### 使用curl

```bash
# 健康检查
curl http://localhost:8080/health

# 获取用户列表
curl http://localhost:8080/api/v1/users

# 创建用户
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John",
    "email": "john@example.com",
    "phone": "13800138000"
  }'

# 获取特定用户
curl http://localhost:8080/api/v1/users/1

# 更新用户
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com"
  }'

# 删除用户
curl -X DELETE http://localhost:8080/api/v1/users/1
```
