package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/models"
	"github.com/liaozzzzzz/code-push-server/internal/service"
	"github.com/liaozzzzzz/code-push-server/internal/utils"
)

// AppController 应用控制器
type AppController struct {
	appService *service.AppService
}

// NewAppController 创建应用控制器实例
func NewAppController() *AppController {
	return &AppController{
		appService: service.NewAppService(),
	}
}

// Create 创建应用
func (c *AppController) Create(ctx *gin.Context) {
	var req models.AppCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	// 从上下文获取用户ID (这里需要中间件设置)
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.HandleAuthError(ctx, models.CodeInvalidToken, "用户未登录")
		return
	}

	app, err := c.appService.Create(userID.(uint), &req)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, app)
}

// GetByID 根据ID获取应用
func (c *AppController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.HandleParamError(ctx, "无效的应用ID")
		return
	}

	app, err := c.appService.GetByID(uint(id))
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, app)
}

// Update 更新应用
func (c *AppController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.HandleParamError(ctx, "无效的应用ID")
		return
	}

	var req models.AppUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	// 从上下文获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.HandleAuthError(ctx, models.CodeInvalidToken, "用户未登录")
		return
	}

	app, err := c.appService.Update(uint(id), userID.(uint), &req)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, app)
}

// Delete 删除应用
func (c *AppController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.HandleParamError(ctx, "无效的应用ID")
		return
	}

	// 从上下文获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.HandleAuthError(ctx, models.CodeInvalidToken, "用户未登录")
		return
	}

	if err := c.appService.Delete(uint(id), userID.(uint)); err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, nil)
}

// List 获取应用列表
func (c *AppController) List(ctx *gin.Context) {
	var pageReq models.PageRequest
	if err := ctx.ShouldBindQuery(&pageReq); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	apps, total, err := c.appService.List(&pageReq)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandlePageSuccess(ctx, apps, pageReq.GetPage(), pageReq.GetSize(), total)
}

// ListByUser 获取用户的应用列表
func (c *AppController) ListByUser(ctx *gin.Context) {
	var pageReq models.PageRequest
	if err := ctx.ShouldBindQuery(&pageReq); err != nil {
		utils.HandleParamError(ctx, "请求参数无效: "+err.Error())
		return
	}

	// 从上下文获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		utils.HandleAuthError(ctx, models.CodeInvalidToken, "用户未登录")
		return
	}

	apps, total, err := c.appService.ListByUser(userID.(uint), &pageReq)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandlePageSuccess(ctx, apps, pageReq.GetPage(), pageReq.GetSize(), total)
}

// GetByBundleID 根据Bundle ID获取应用
func (c *AppController) GetByBundleID(ctx *gin.Context) {
	bundleID := ctx.Param("bundle_id")
	if bundleID == "" {
		utils.HandleParamError(ctx, "Bundle ID不能为空")
		return
	}

	app, err := c.appService.GetByBundleID(bundleID)
	if err != nil {
		if bizErr, ok := err.(*models.BusinessError); ok {
			utils.HandleBusinessError(ctx, bizErr.GetCode(), bizErr.Error())
		} else {
			utils.HandleBusinessError(ctx, models.CodeServiceError, "服务错误")
		}
		return
	}

	utils.HandleSuccess(ctx, app)
}
