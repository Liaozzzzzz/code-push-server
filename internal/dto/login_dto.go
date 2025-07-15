package dto

import "github.com/liaozzzzzz/code-push-server/internal/entity"

// LoginForm 登录表单
type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResult 登录结果
type LoginResult struct {
	Token string      `json:"token"`
	User  entity.User `json:"user"`
}
