package dto

import (
	"time"

	"github.com/liaozzzzzz/code-push-server/internal/entity"
	"github.com/liaozzzzzz/code-push-server/internal/types"
)

// MenuCreateRequest 创建菜单请求
type MenuCreateRequest struct {
	MenuName string           `json:"menuName" binding:"required,min=3,max=50"`
	MenuKey  string           `json:"menuKey" binding:"required,min=3,max=50"`
	MenuSort int32            `json:"menuSort" binding:"required,min=0"`
	Status   types.MenuStatus `json:"status" binding:"required,oneof='1' '0'"`
	Remark   string           `json:"remark" binding:"omitempty,max=255"`
}

// MenuUpdateRequest 更新菜单请求
type MenuUpdateRequest struct {
	Id       string           `json:"id" binding:"required"`
	MenuName string           `json:"menuName" binding:"omitempty,min=3,max=50"`
	MenuSort int32            `json:"menuSort" binding:"omitempty,min=0"`
	Status   types.MenuStatus `json:"status" binding:"required,oneof='1' '0'"`
	Remark   string           `json:"remark" binding:"omitempty,max=255"`
}

// MenuResponse 菜单响应
type MenuResponse struct {
	MenuID     int32            `json:"menuId"`
	MenuName   string           `json:"menuName"`
	MenuSort   int32            `json:"menuSort"`
	MenuStatus types.MenuStatus `json:"menuStatus"`
	Remark     string           `json:"remark"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}

// ToMenuResponse 将菜单实体转换为响应DTO
func ToMenuResponse(menu *entity.Menu) *MenuResponse {
	return &MenuResponse{
		MenuID:     menu.MenuID,
		MenuName:   menu.MenuName,
		MenuSort:   menu.MenuSort,
		MenuStatus: menu.MenuStatus,
		Remark:     menu.Remark,
		CreatedAt:  menu.CreatedAt,
		UpdatedAt:  menu.UpdatedAt,
	}
}
