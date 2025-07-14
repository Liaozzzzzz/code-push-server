package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/utils/errors"
	"github.com/liaozzzzzz/code-push-server/internal/utils/response"
)

// AuthMiddleware JWT认证中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "缺少认证头部"))
			c.Abort()
			return
		}

		// 检查Bearer格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "无效的认证格式"))
			c.Abort()
			return
		}

		// 获取token
		token := parts[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "缺少认证令牌"))
			c.Abort()
			return
		}

		// 验证token（这里简化处理，实际应该验证JWT）
		// 在实际项目中，这里应该：
		// 1. 解析JWT token
		// 2. 验证token的有效性
		// 3. 获取用户信息
		// 4. 设置用户上下文

		// 简单的token验证示例
		if !strings.HasPrefix(token, "mock_token_") {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "无效的认证令牌"))
			c.Abort()
			return
		}

		// 设置用户信息到上下文（示例）
		username := strings.TrimPrefix(token, "mock_token_")
		c.Set("username", username)
		c.Set("user_id", 1) // 示例用户ID

		c.Next()
	}
}
