package entity

import (
	"time"

	"gorm.io/gorm"
)

// RoleMenu 角色菜单关联实体模型
type RoleMenu struct {
	RoleID    int32          `json:"roleId" gorm:"type:bigint(20);primaryKey;comment:'角色ID'"`
	MenuID    int32          `json:"menuId" gorm:"type:bigint(20);primaryKey;comment:'菜单ID'"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:'创建时间'"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:'更新时间'"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:'删除时间'"`
}

// TableName 指定表名
func (RoleMenu) TableName() string {
	return "role_menu"
}
