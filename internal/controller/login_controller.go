package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/dto"
	"github.com/liaozzzzzz/code-push-server/internal/service"
	"github.com/liaozzzzzz/code-push-server/internal/utils/response"
)

// LoginController 登录控制器
type LoginController struct {
	loginService *service.LoginService
}

// NewLoginController 创建登录控制器实例
func NewLoginController() *LoginController {
	return &LoginController{
		loginService: service.NewLoginService(),
	}
}

// @Summary      账户密码登录
// @Description  账户密码登录
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        object		body      dto.LoginForm  true  "Login Form"
// @Success      200		{object}  response.Response{data=dto.LoginResult}
// @Router       /login		[post]
func (c *LoginController) Login(ctx *gin.Context) {
	data := new(dto.LoginForm)
	if err := response.ParseJSON(ctx, data); err != nil {
		response.HandleError(ctx, err)
		return
	}

	result, err := c.loginService.Login(data)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, result)
}

// Logout 用户登出
func (c *LoginController) Logout(ctx *gin.Context) {
	// 简单的登出处理，实际应该清除token等
	response.HandleSuccess(ctx, gin.H{"message": "登出成功"})
}
