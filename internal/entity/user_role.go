package entity

import (
	"time"

	"gorm.io/gorm"
)

// UserRole 用户角色关联实体模型
type UserRole struct {
	UserID    int32          `json:"userId" gorm:"type:bigint(20);primaryKey;comment:'用户ID'"`
	RoleID    int32          `json:"roleId" gorm:"type:bigint(20);primaryKey;comment:'角色ID'"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:'创建时间'"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:'更新时间'"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:'删除时间'"`
}

// TableName 指定表名
func (UserRole) TableName() string {
	return "user_role"
}
