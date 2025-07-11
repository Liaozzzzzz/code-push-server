package models

import (
	"time"

	"gorm.io/gorm"
)

// App 应用模型
type App struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Platform    string         `json:"platform" gorm:"not null"` // ios, android
	BundleID    string         `json:"bundle_id" gorm:"not null"`
	Description string         `json:"description"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (App) TableName() string {
	return "apps"
}

// AppCreateRequest 创建应用请求
type AppCreateRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=100"`
	Platform    string `json:"platform" binding:"required,oneof=ios android"`
	BundleID    string `json:"bundle_id" binding:"required"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

// AppUpdateRequest 更新应用请求
type AppUpdateRequest struct {
	Name        string `json:"name" binding:"omitempty,min=1,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

// AppResponse 应用响应
type AppResponse struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Platform    string       `json:"platform"`
	BundleID    string       `json:"bundle_id"`
	Description string       `json:"description"`
	User        UserResponse `json:"user"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// ToResponse 转换为响应格式
func (a *App) ToResponse() *AppResponse {
	return &AppResponse{
		ID:          a.ID,
		Name:        a.Name,
		Platform:    a.Platform,
		BundleID:    a.BundleID,
		Description: a.Description,
		User:        *a.User.ToResponse(),
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}
