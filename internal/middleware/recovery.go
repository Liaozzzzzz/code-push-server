package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/models"
)

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, models.Error(models.CodeInternalError, "服务器内部错误: "+err))
		} else {
			c.JSON(http.StatusInternalServerError, models.Error(models.CodeInternalError, "服务器内部错误"))
		}
		c.Abort()
	})
}
