package entity

import (
	"time"

	"github.com/liaozzzzzz/code-push-server/internal/types"
	"gorm.io/gorm"
)

// Menu 菜单实体模型
type Menu struct {
	MenuID      int32             `json:"menuId" gorm:"type:bigint(20);primaryKey;autoIncrement;comment:'菜单ID'"`
	MenuName    string            `json:"menuName" gorm:"type:varchar(50);not null;uniqueIndex;comment:'菜单名称'"`
	ParentID    int32             `json:"parentId" gorm:"type:bigint(20);default:0;comment:'父菜单ID'"`
	Perms       string            `json:"perms" gorm:"type:varchar(255);comment:'权限标识'"`
	MenuType    types.MenuType    `json:"menuType" gorm:"type:enum('1','2','3');default:'2';comment:'菜单类型'"`
	MenuVisible types.MenuVisible `json:"menuVisible" gorm:"type:enum('1','0');default:'1';comment:'是否显示'"`
	MenuIsLink  types.MenuIsLink  `json:"menuIsLink" gorm:"type:enum('1','0');default:'0';comment:'是否外链'"`
	Icon        string            `json:"icon" gorm:"type:varchar(50);comment:'菜单图标'"`
	Path        string            `json:"path" gorm:"type:varchar(255);comment:'菜单路径'"`
	MenuSort    int32             `json:"sort" gorm:"type:int(11);not null;default:0;comment:'菜单排序'"`
	MenuStatus  types.MenuStatus  `json:"status" gorm:"type:enum('1','0');default:'1';comment:'状态'"`
	Remark      string            `json:"remark" gorm:"type:varchar(255);comment:'备注'"`
	CreatedAt   time.Time         `json:"createdAt" gorm:"comment:'创建时间'"`
	UpdatedAt   time.Time         `json:"updatedAt" gorm:"comment:'更新时间'"`
	DeletedAt   gorm.DeletedAt    `json:"-" gorm:"index;comment:'删除时间'"`
}

// TableName 指定表名
func (Menu) TableName() string {
	return "menus"
}
