package dto

import (
	"time"

	"github.com/liaozzzzzz/code-push-server/internal/entity"
)

// RoleMenuCreateRequest 创建角色菜单关联请求
type RoleMenuCreateRequest struct {
	RoleID int32 `json:"roleId" binding:"required"`
	MenuID int32 `json:"menuId" binding:"required"`
}

// RoleMenuUpdateRequest 更新角色菜单关联请求
type RoleMenuUpdateRequest struct {
	RoleID int32 `json:"roleId" binding:"required"`
	MenuID int32 `json:"menuId" binding:"required"`
}

// RoleMenuResponse 角色菜单关联响应
type RoleMenuResponse struct {
	RoleID    int32     `json:"roleId"`
	MenuID    int32     `json:"menuId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToRoleMenuResponse 将角色菜单关联实体转换为响应DTO
func ToRoleMenuResponse(roleMenu *entity.RoleMenu) *RoleMenuResponse {
	return &RoleMenuResponse{
		RoleID:    roleMenu.RoleID,
		MenuID:    roleMenu.MenuID,
		CreatedAt: roleMenu.CreatedAt,
		UpdatedAt: roleMenu.UpdatedAt,
	}
}
