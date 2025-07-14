package entity

import (
	"time"

	"gorm.io/gorm"
)

type RoleStatus string

const (
	// 启用
	RoleEnabled RoleStatus = "1"
	// 禁用
	RoleDisabled RoleStatus = "0"
)

// Role 角色实体模型
type Role struct {
	RoleID     int32          `json:"roleId" gorm:"type:bigint(20);primaryKey;autoIncrement;comment:'角色ID'"`
	RoleName   string         `json:"roleName" gorm:"type:varchar(50);not null;uniqueIndex;comment:'角色名称'"`
	RoleKey    string         `json:"roleKey" gorm:"type:varchar(50);not null;uniqueIndex;comment:'角色键'"`
	RoleSort   int32          `json:"roleSort" gorm:"type:int(11);not null;default:0;comment:'角色排序'"`
	RoleStatus RoleStatus     `json:"roleStatus" gorm:"type:enum('1','0');default:'1';comment:'状态'"`
	Remark     string         `json:"remark" gorm:"type:varchar(255);comment:'备注'"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"comment:'创建时间'"`
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"comment:'更新时间'"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index;comment:'删除时间'"`
}

// TableName 指定表名
func (Role) TableName() string {
	return "roles"
}
