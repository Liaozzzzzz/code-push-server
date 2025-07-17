package dto

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/liaozzzzzz/code-push-server/internal/entity"
	"github.com/liaozzzzzz/code-push-server/internal/types"
)

// MenuResponse 菜单响应
type MenuResponse struct {
	MenuID      int64             `json:"menuId"`
	MenuName    string            `json:"menuName"`
	ParentID    int64             `json:"parentId"`
	Perms       *string           `json:"perms"`
	MenuType    types.MenuType    `json:"menuType"`
	MenuVisible types.MenuVisible `json:"menuVisible"`
	MenuIsLink  types.MenuIsLink  `json:"menuIsLink"`
	Icon        *string           `json:"icon"`
	Path        *string           `json:"path"`
	Sort        int32             `json:"sort"`
	Status      types.MenuStatus  `json:"status"`
	Remark      *string           `json:"remark"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

// MenuTreeResponse 菜单树响应
type MenuTreeResponse struct {
	MenuResponse
	Children []*MenuTreeResponse `json:"children"`
}

// MenuCreateForm 创建菜单表单
type MenuCreateForm struct {
	MenuName    string            `json:"menuName" binding:"required,min=2,max=50"`
	ParentID    *int64            `json:"parentId" binding:"required,min=0"`
	Perms       string            `json:"perms" binding:"omitempty,max=255"`
	MenuType    types.MenuType    `json:"menuType" binding:"required,oneof='1' '2' '3'"`
	MenuVisible types.MenuVisible `json:"menuVisible" binding:"required,oneof='1' '0'"`
	MenuIsLink  types.MenuIsLink  `json:"menuIsLink" binding:"required,oneof='1' '0'"`
	Icon        *string           `json:"icon" binding:"omitempty,max=50"`
	Path        *string           `json:"path" binding:"omitempty,max=255"`
	Sort        int32             `json:"sort" binding:"required,min=0"`
	Status      types.MenuStatus  `json:"status" binding:"required,oneof='1' '0'"`
	Remark      *string           `json:"remark" binding:"omitempty,max=255"`
}

// MenuUpdateForm 更新菜单表单
type MenuUpdateForm struct {
	MenuCreateForm
	MenuID int64 `json:"menuId" binding:"required"`
}

// MenuDeleteForm 删除菜单表单
type MenuDeleteForm struct {
	MenuID int64 `json:"menuId" binding:"required"`
}

func BuildMenuTree(menuList []*entity.Menu, parentID int64) []*MenuTreeResponse {
	menuTree := make([]*MenuTreeResponse, 0)
	for _, menu := range menuList {
		if menu.ParentID != parentID {
			continue
		}
		var menuResponse MenuTreeResponse
		if err := copier.Copy(&menuResponse, menu); err != nil {
			continue
		}
		menuResponse.Children = BuildMenuTree(menuList, menu.MenuID)
		menuTree = append(menuTree, &menuResponse)
	}
	return menuTree
}
