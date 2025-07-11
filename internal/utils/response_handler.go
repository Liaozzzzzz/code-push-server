package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/models"
)

// HandleParamError 处理参数错误
func HandleParamError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, models.Error(models.CodeInvalidParams, message))
}

// HandleSuccess 处理成功响应
func HandleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Success(data))
}

// HandlePageSuccess 处理分页成功响应
func HandlePageSuccess(c *gin.Context, data interface{}, page, size int, total int64) {
	c.JSON(http.StatusOK, models.PageSuccess(data, page, size, total))
}

// HandleAuthError 处理认证错误（返回401状态码）
func HandleAuthError(c *gin.Context, code int, message string) {
	c.JSON(http.StatusUnauthorized, models.Error(code, message))
	c.Abort()
}

// HandleInternalError 处理内部服务器错误（返回500状态码）
func HandleInternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, models.Error(models.CodeInternalError, message))
	c.Abort()
}

// HandleBusinessError 处理业务错误（返回200状态码）
func HandleBusinessError(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, models.Error(code, message))
}
