# Code Push Server API 文档

## 概述

这是一个基于 Go 语言和 Gin 框架构建的代码推送服务器，采用分层架构设计，包含 Controller、Service、DAO 三层。

## 架构设计

### 分层架构

```
┌─────────────────────────────────────────────────────────┐
│                    Controller Layer                     │
│                   (HTTP 请求处理)                       │
├─────────────────────────────────────────────────────────┤
│                    Service Layer                        │
│                   (业务逻辑处理)                         │
├─────────────────────────────────────────────────────────┤
│                      DAO Layer                          │
│                   (数据访问对象)                         │
├─────────────────────────────────────────────────────────┤
│                     Database                            │
│                   (SQLite/MySQL)                        │
└─────────────────────────────────────────────────────────┘
```

### 目录结构

```
internal/
├── controller/          # 控制器层
│   ├── user_controller.go
│   └── app_controller.go
├── service/             # 服务层
│   ├── user_service.go
│   └── app_service.go
├── dao/                 # 数据访问层
│   ├── user_dao.go
│   └── app_dao.go
├── models/              # 数据模型
│   ├── user.go
│   ├── app.go
│   └── response.go
├── middleware/          # 中间件
│   ├── auth.go
│   ├── cors.go
│   ├── logger.go
│   └── recovery.go
├── router/              # 路由配置
│   └── router.go
├── server/              # HTTP服务器
│   └── server.go
├── database/            # 数据库连接
│   └── database.go
├── utils/               # 工具函数
│   └── error_handler.go
└── config/              # 配置管理
    └── config.go
```

## API 端点

### 基础信息

- **基础 URL**: `http://localhost:8040`
- **内容类型**: `application/json`
- **认证方式**: Bearer Token

### 状态码设计

#### HTTP 状态码

HTTP 状态码只用于表示 HTTP 层面的状态：

- `200` - 请求成功（业务成功或失败都返回 200）
- `401` - 未授权（认证失败）
- `500` - 服务器内部错误（系统级异常）

#### 业务状态码

业务状态码用于表示具体的业务处理结果：

```json
{
  "code": 0,           // 业务状态码
  "message": "操作成功", // 状态描述
  "data": {...}        // 响应数据
}
```

**业务状态码定义：**

- `0` - 操作成功
- `1001` - 参数错误
- `1002` - 资源已存在
- `1003` - 资源不存在
- `1004` - 权限不足
- `1005` - 登录失败
- `1006` - 账户被禁用
- `1007` - 无效的令牌
- `1008` - 令牌过期
- `2001` - 内部服务器错误
- `2002` - 数据库错误
- `2003` - 服务错误

### 健康检查

#### GET /health

检查服务器状态

**响应示例:**

```json
{
  "status": "ok",
  "message": "服务运行正常"
}
```

### 用户管理

#### POST /api/auth/login

用户登录

**请求体:**

```json
{
  "username": "admin",
  "password": "password123"
}
```

**响应示例:**

成功登录：

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "role": "admin",
    "status": "active",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

登录失败：

```json
{
  "code": 1005,
  "message": "登录失败"
}
```

#### POST /api/users

创建用户

**请求体:**

```json
{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "password123",
  "role": "developer"
}
```

**响应示例:**

成功：

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "id": 2,
    "username": "newuser",
    "email": "newuser@example.com",
    "role": "developer",
    "status": "active",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

用户名已存在：

```json
{
  "code": 1002,
  "message": "用户名已存在"
}
```

#### GET /api/users

获取用户列表

**查询参数:**

- `page` (int, 可选): 页码，默认为 1
- `size` (int, 可选): 每页大小，默认为 10

**响应示例:**

```json
{
  "code": 0,
  "message": "操作成功",
  "data": [
    {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "admin",
      "status": "active",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "page": {
    "current": 1,
    "size": 10,
    "total": 1,
    "total_pages": 1
  }
}
```

#### GET /api/users/{id}

获取用户详情

**路径参数:**

- `id` (int): 用户 ID

**响应示例:**

成功：

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "role": "admin",
    "status": "active",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

用户不存在：

```json
{
  "code": 1003,
  "message": "用户不存在"
}
```

#### PUT /api/users/{id}

更新用户

**请求体:**

```json
{
  "username": "updateduser",
  "email": "updated@example.com",
  "role": "developer",
  "status": "active"
}
```

#### DELETE /api/users/{id}

删除用户

#### PUT /api/users/{id}/status

更新用户状态

**请求体:**

```json
{
  "status": "inactive"
}
```

### 应用管理

#### GET /api/apps

获取应用列表

**查询参数:**

- `page` (int, 可选): 页码，默认为 1
- `size` (int, 可选): 每页大小，默认为 10

#### GET /api/apps/{id}

获取应用详情

**路径参数:**

- `id` (int): 应用 ID

#### GET /api/apps/bundle/{bundle_id}

根据 Bundle ID 获取应用

**路径参数:**

- `bundle_id` (string): Bundle ID

### 需要认证的端点

以下端点需要在请求头中包含 `Authorization: Bearer <token>`

认证失败时返回 HTTP 401 状态码：

```json
{
  "code": 1007,
  "message": "无效的令牌"
}
```

#### POST /api/apps

创建应用

**请求体:**

```json
{
  "name": "My App",
  "platform": "ios",
  "bundle_id": "com.example.myapp",
  "description": "这是我的应用"
}
```

**响应示例:**

成功：

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "id": 1,
    "name": "My App",
    "platform": "ios",
    "bundle_id": "com.example.myapp",
    "description": "这是我的应用",
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "admin",
      "status": "active",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

Bundle ID 已存在：

```json
{
  "code": 1002,
  "message": "Bundle ID已存在"
}
```

#### PUT /api/apps/{id}

更新应用

**请求体:**

```json
{
  "name": "Updated App Name",
  "description": "更新后的描述"
}
```

权限不足：

```json
{
  "code": 1004,
  "message": "权限不足"
}
```

#### DELETE /api/apps/{id}

删除应用

#### GET /api/my/apps

获取当前用户的应用列表

## 响应格式

### 成功响应

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {...}
}
```

### 分页响应

```json
{
  "code": 0,
  "message": "操作成功",
  "data": [...],
  "page": {
    "current": 1,
    "size": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

### 错误响应

```json
{
  "code": 1001,
  "message": "参数错误"
}
```

## 状态码

### HTTP 状态码

- `200` - 成功（包含业务成功和失败）
- `401` - 未授权（认证失败）
- `500` - 服务器内部错误（系统异常）

### 业务状态码

- `0` - 操作成功
- `1001` - 参数错误
- `1002` - 资源已存在
- `1003` - 资源不存在
- `1004` - 权限不足
- `1005` - 登录失败
- `1006` - 账户被禁用
- `1007` - 无效的令牌
- `1008` - 令牌过期
- `2001` - 内部服务器错误
- `2002` - 数据库错误
- `2003` - 服务错误

## 中间件

### CORS

支持跨域请求，可在配置文件中设置允许的域名、方法和头部。

### 日志

记录所有 HTTP 请求的详细信息，包括 IP、方法、路径、状态码和响应时间。

### 认证

验证 Bearer Token，提取用户信息并存储到请求上下文中。认证失败时返回 HTTP 401 状态码。

### 恢复

捕获 panic 异常，返回 HTTP 500 状态码和友好的错误信息。

## 数据库模型

### User (用户)

```go
type User struct {
    ID        uint      `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    Role      string    `json:"role"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### App (应用)

```go
type App struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Platform    string    `json:"platform"`
    BundleID    string    `json:"bundle_id"`
    Description string    `json:"description"`
    UserID      uint      `json:"user_id"`
    User        User      `json:"user"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## 配置

服务器配置通过 TOML 文件管理，支持多环境配置（开发、测试、生产）。

主要配置项：

- HTTP 服务器设置
- 数据库连接
- 安全设置（JWT、CORS）
- 日志配置
- 缓存配置

## 启动服务器

```bash
go run main.go start --config ./configs --env dev
```

服务器将在配置的端口上启动（默认 8040）。
