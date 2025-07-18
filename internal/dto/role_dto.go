package dto

import (
	"time"

	"github.com/liaozzzzzz/code-push-server/internal/types"
)

// RoleCreateRequest 创建角色请求
type RoleCreateForm struct {
	RoleName string           `json:"roleName" binding:"required,min=3,max=50"`
	RoleKey  string           `json:"roleKey" binding:"required,min=3,max=50"`
	Status   types.RoleStatus `json:"status" binding:"required,oneof='1' '0'"`
	Remark   *string          `json:"remark" binding:"omitempty,max=255"`
}

// RoleUpdateForm 更新角色请求
type RoleUpdateForm struct {
	RoleCreateForm
	RoleID int64 `json:"roleId" binding:"required"`
}

// RoleResponse 角色响应
type RoleResponse struct {
	RoleID    int64            `json:"roleId"`
	RoleName  string           `json:"roleName"`
	RoleKey   string           `json:"roleKey"`
	Status    types.RoleStatus `json:"status"`
	Remark    *string          `json:"remark"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
