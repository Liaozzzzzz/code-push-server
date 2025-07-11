package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/models"
	"github.com/liaozzzzzz/code-push-server/internal/service"
	"github.com/liaozzzzzz/code-push-server/internal/utils"
)

// UserController 用户控制器
type UserController struct {
	userService *service.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// Create 创建用户
func (c *UserController) Create(ctx *gin.Context) {
	var req models.UserCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	user, err := c.userService.Create(&req)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, user)
}

// GetByID 根据ID获取用户
func (c *UserController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.HandleParamError(ctx, "无效的用户ID")
		return
	}

	user, err := c.userService.GetByID(uint(id))
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, user)
}

// Update 更新用户
func (c *UserController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.HandleParamError(ctx, "无效的用户ID")
		return
	}

	var req models.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	user, err := c.userService.Update(uint(id), &req)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, user)
}

// Delete 删除用户
func (c *UserController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.HandleParamError(ctx, "无效的用户ID")
		return
	}

	if err := c.userService.Delete(uint(id)); err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, nil)
}

// List 获取用户列表
func (c *UserController) List(ctx *gin.Context) {
	var pageReq models.PageRequest
	if err := ctx.ShouldBindQuery(&pageReq); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	users, total, err := c.userService.List(&pageReq)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandlePageSuccess(ctx, users, pageReq.GetPage(), pageReq.GetSize(), total)
}

// Login 用户登录
func (c *UserController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	user, err := c.userService.Login(req.Username, req.Password)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, user)
}

// UpdateStatus 更新用户状态
func (c *UserController) UpdateStatus(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.HandleParamError(ctx, "无效的用户ID")
		return
	}

	var req StatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	if err := c.userService.UpdateStatus(uint(id), req.Status); err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, nil)
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// StatusRequest 状态请求
type StatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active inactive"`
}
