package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UserStatus string

const (
	// 启用
	UserEnabled UserStatus = "1"
	// 禁用
	UserDisabled UserStatus = "0"
)

// User 用户实体模型
type User struct {
	UserID     int32          `json:"userId" gorm:"type:bigint(20);primaryKey;autoIncrement;comment:'用户ID'"`
	Username   string         `json:"username" gorm:"type:varchar(50);not null;uniqueIndex;comment:'用户名'"`
	Nickname   sql.NullString `json:"nickname" gorm:"type:varchar(50);comment:'昵称'"`
	Email      string         `json:"email" gorm:"type:varchar(100);not null;uniqueIndex;comment:'邮箱'"`
	Password   string         `json:"-" gorm:"type:varchar(255);not null;comment:'密码'"`
	Avatar     sql.NullString `json:"avatar" gorm:"type:varchar(255);default:null;comment:'头像'"`
	AckCode    sql.NullString `json:"ackCode" gorm:"type:varchar(255);default:null;comment:'验证码'"`
	UserStatus UserStatus     `json:"userStatus" gorm:"type:enum('1','0');default:'1';comment:'状态'"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"comment:'创建时间'"`
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"comment:'更新时间'"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index;comment:'删除时间'"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
