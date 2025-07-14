package dto

import (
	"time"

	"github.com/liaozzzzzz/code-push-server/internal/entity"
)

// UserRoleCreateRequest 创建用户角色关联请求
type UserRoleCreateRequest struct {
	UserID int32 `json:"userId" binding:"required"`
	RoleID int32 `json:"roleId" binding:"required"`
}

// UserRoleUpdateRequest 更新用户角色关联请求
type UserRoleUpdateRequest struct {
	UserID int32 `json:"userId" binding:"required"`
	RoleID int32 `json:"roleId" binding:"required"`
}

// UserRoleResponse 用户角色关联响应
type UserRoleResponse struct {
	UserID    int32     `json:"userId"`
	RoleID    int32     `json:"roleId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToUserRoleResponse 将用户角色关联实体转换为响应DTO
func ToUserRoleResponse(userRole *entity.UserRole) *UserRoleResponse {
	return &UserRoleResponse{
		UserID:    userRole.UserID,
		RoleID:    userRole.RoleID,
		CreatedAt: userRole.CreatedAt,
		UpdatedAt: userRole.UpdatedAt,
	}
}
