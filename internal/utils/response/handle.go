package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/utils/errors"
	"gorm.io/gorm"
)

// HandleSuccess 处理成功响应
func HandleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Success(data))
}

// HandlePageSuccess 处理分页成功响应
func HandlePageSuccess(c *gin.Context, data interface{}, page, size int, total int) {
	c.JSON(http.StatusOK, PageSuccess(data, page, size, total))
}

// HandleParamError 处理参数错误
func HandleParamError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Error(errors.CodeInvalidParams, message))
}

// HandleAuthError 处理认证错误（返回401状态码）
func HandleAuthError(c *gin.Context, code errors.BusinessCode, message string) {
	c.JSON(http.StatusUnauthorized, Error(code, message))
	c.Abort()
}

// HandleInternalError 处理内部服务器错误（返回500状态码）
func HandleInternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Error(errors.CodeInternalError, message))
	c.Abort()
}

// AbortWithError 处理错误并返回响应
func HandleError(c *gin.Context, err error) {
	if c.Writer.Written() {
		return
	}

	switch e := err.(type) {
	case *errors.BusinessError:
		c.JSON(http.StatusOK, Error(e.GetCode(), e.GetMessage()))
	default:
		// 检查是否是GORM记录未找到错误
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, Error(errors.CodeResourceNotFound, "资源不存在"))
			return
		}
		// 其他未知错误
		c.JSON(http.StatusOK, Error(errors.CodeInternalError, "服务器内部错误"))
	}
}
