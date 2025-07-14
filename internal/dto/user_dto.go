package dto

import (
	"time"

	"github.com/liaozzzzzz/code-push-server/internal/entity"
)

// UserCreateRequest 用户创建请求
type UserCreateRequest struct {
	Username   string            `json:"username" binding:"required,min=3,max=50"`
	Nickname   string            `json:"nickname" binding:"omitempty,min=3,max=50"`
	Avatar     string            `json:"avatar" binding:"max=255"`
	Email      string            `json:"email" binding:"required,email"`
	Password   string            `json:"password" binding:"required,min=8,max=255"`
	UserStatus entity.UserStatus `json:"userStatus" binding:"required,oneof='1' '0'"`
}

// UserUpdateRequest 用户更新请求
type UserUpdateRequest struct {
	Id       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
	Email    string `json:"email" binding:"omitempty,email"`
}

// UserResponse 用户响应
type UserResponse struct {
	UserID    int32     `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToUserResponse 将用户实体转换为响应DTO
func ToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		UserID:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
