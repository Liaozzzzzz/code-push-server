# 代码重构指南

## 问题分析

当前 `internal/models/` 目录存在职责混合的问题：

- 数据库模型（`user.go`, `menu.go`, `role.go` 等）
- 业务错误处理（`error.go`, `codes.go`）
- HTTP 响应格式（`response.go`）
- 数据传输对象（各种 Request/Response 结构）

这违反了单一职责原则，使代码难以维护和理解。

## 新的架构设计

### 📁 目录结构

```
internal/
├── entity/          # 数据库实体模型
│   ├── user.go      # 用户实体
│   ├── menu.go      # 菜单实体
│   ├── role.go      # 角色实体
│   └── ...
├── dto/             # 数据传输对象
│   ├── user_dto.go  # 用户相关的请求/响应
│   ├── menu_dto.go  # 菜单相关的请求/响应
│   ├── login_dto.go # 登录相关的请求/响应
│   └── ...
├── pkg/             # 通用包
│   ├── response/    # HTTP响应处理
│   │   └── response.go
│   └── errors/      # 错误处理
│       ├── codes.go
│       └── errors.go
└── models/          # 旧目录（可删除）
```

### 🎯 各目录职责

#### `internal/entity/`

- **职责**: 数据库实体模型
- **包含**:
  - GORM 模型定义
  - 数据库表结构
  - 基本的数据验证
- **不包含**:
  - 业务逻辑
  - HTTP 请求/响应结构
  - 错误处理

#### `internal/dto/`

- **职责**: 数据传输对象
- **包含**:
  - HTTP 请求结构（CreateRequest, UpdateRequest 等）
  - HTTP 响应结构（Response 等）
  - 数据转换方法（Entity -> DTO）
- **不包含**:
  - 数据库相关的标签
  - 业务逻辑

#### `internal/pkg/errors/`

- **职责**: 错误处理
- **包含**:
  - 业务状态码定义
  - 业务错误结构
  - 错误创建和转换方法
- **特点**: 通用的、可复用的错误处理

#### `internal/pkg/response/`

- **职责**: HTTP 响应格式
- **包含**:
  - 统一的响应结构
  - 分页响应结构
  - 响应创建方法
- **特点**: 通用的、可复用的响应处理

## 迁移步骤

### 1. 迁移数据库模型

```bash
# 将数据库模型移动到entity目录
mv internal/models/user.go internal/entity/user.go
mv internal/models/menu.go internal/entity/menu.go
mv internal/models/role.go internal/entity/role.go
# ... 其他模型文件
```

### 2. 提取 DTO

从原始模型文件中提取请求/响应结构到 dto 目录：

```go
// 从 internal/models/user.go 提取
type UserCreateRequest struct { ... }
type UserUpdateRequest struct { ... }
type UserResponse struct { ... }

// 移动到 internal/dto/user_dto.go
```

### 3. 更新导入路径

```go
// 旧的导入
import "github.com/liaozzzzzz/code-push-server/internal/models"

// 新的导入
import (
    "github.com/liaozzzzzz/code-push-server/internal/entity"
    "github.com/liaozzzzzz/code-push-server/internal/dto"
    "github.com/liaozzzzzz/code-push-server/internal/pkg/errors"
    "github.com/liaozzzzzz/code-push-server/internal/pkg/response"
)
```

### 4. 更新代码使用

```go
// 旧的使用方式
user := &models.User{}
response := models.Success(data)
err := models.NewBusinessError(models.CodeInvalidParams, "参数错误")

// 新的使用方式
user := &entity.User{}
response := response.Success(data)
err := errors.NewBusinessError(errors.CodeInvalidParams, "参数错误")
```

## 优势

### ✅ 单一职责

- 每个包都有明确的职责
- 代码更加模块化
- 易于理解和维护

### ✅ 可复用性

- 错误处理和响应格式可以在多个模块中复用
- 减少代码重复

### ✅ 可测试性

- 每个包可以独立测试
- 依赖关系更加清晰

### ✅ 扩展性

- 新增功能时更容易找到合适的位置
- 修改影响范围更小

## 注意事项

1. **渐进式迁移**: 可以逐步迁移，不需要一次性完成
2. **测试覆盖**: 迁移后要确保测试用例仍然通过
3. **文档更新**: 更新相关文档和 API 文档
4. **团队沟通**: 确保团队成员了解新的架构

## 示例代码

### Entity 示例

```go
// internal/entity/user.go
package entity

type User struct {
    UserID    int32     `gorm:"primaryKey"`
    Username  string    `gorm:"uniqueIndex"`
    Email     string    `gorm:"uniqueIndex"`
    // ... 其他字段
}

func (User) TableName() string {
    return "users"
}
```

### DTO 示例

```go
// internal/dto/user_dto.go
package dto

type UserCreateRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
}

type UserResponse struct {
    UserID   int32  `json:"userId"`
    Username string `json:"username"`
    Email    string `json:"email"`
}
```

### 错误处理示例

```go
// internal/pkg/errors/errors.go
package errors

func NewBusinessError(code BusinessCode, message string) *BusinessError {
    return &BusinessError{
        Code:    code,
        Message: message,
    }
}
```

### 响应处理示例

```go
// internal/pkg/response/response.go
package response

func Success(data interface{}) *Response {
    return &Response{
        Code:    10000,
        Message: "操作成功",
        Data:    data,
    }
}
```
