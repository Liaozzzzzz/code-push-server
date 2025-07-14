package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/utils/errors"
	"github.com/liaozzzzzz/code-push-server/internal/utils/response"
)

// RecoveryMiddleware 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录错误日志
				fmt.Printf("Panic recovered: %v\n", err)

				// 返回统一的错误响应
				if errStr, ok := err.(string); ok {
					c.JSON(http.StatusInternalServerError, response.Error(errors.CodeInternalError, "服务器内部错误: "+errStr))
				} else {
					c.JSON(http.StatusInternalServerError, response.Error(errors.CodeInternalError, "服务器内部错误"))
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
