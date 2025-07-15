package dto

import (
	"time"

	"github.com/liaozzzzzz/code-push-server/internal/entity"
	"github.com/liaozzzzzz/code-push-server/internal/types"
)

// RoleCreateRequest 创建角色请求
type RoleCreateRequest struct {
	RoleName string           `json:"roleName" binding:"required,min=3,max=50"`
	RoleKey  string           `json:"roleKey" binding:"required,min=3,max=50"`
	RoleSort int32            `json:"roleSort" binding:"required,min=0"`
	Status   types.RoleStatus `json:"status" binding:"required,oneof='1' '0'"`
	Remark   string           `json:"remark" binding:"omitempty,max=255"`
}

// RoleUpdateRequest 更新角色请求
type RoleUpdateRequest struct {
	Id         string           `json:"id" binding:"required"`
	RoleName   string           `json:"roleName" binding:"omitempty,min=3,max=50"`
	RoleKey    string           `json:"roleKey" binding:"omitempty,min=3,max=50"`
	RoleSort   int32            `json:"roleSort" binding:"omitempty,min=0"`
	RoleStatus types.RoleStatus `json:"roleStatus" binding:"required,oneof='1' '0'"`
	Remark     string           `json:"remark" binding:"omitempty,max=255"`
}

// RoleResponse 角色响应
type RoleResponse struct {
	RoleID     int32            `json:"roleId"`
	RoleName   string           `json:"roleName"`
	RoleKey    string           `json:"roleKey"`
	RoleSort   int32            `json:"roleSort"`
	RoleStatus types.RoleStatus `json:"roleStatus"`
	Remark     string           `json:"remark"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}

// ToRoleResponse 将角色实体转换为响应DTO
func ToRoleResponse(role *entity.Role) *RoleResponse {
	return &RoleResponse{
		RoleID:     role.RoleID,
		RoleName:   role.RoleName,
		RoleKey:    role.RoleKey,
		RoleSort:   role.RoleSort,
		RoleStatus: role.RoleStatus,
		Remark:     role.Remark,
		CreatedAt:  role.CreatedAt,
		UpdatedAt:  role.UpdatedAt,
	}
}
