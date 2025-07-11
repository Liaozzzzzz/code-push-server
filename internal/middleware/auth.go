package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/models"
)

// AuthRequired 认证中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头部
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.Error(models.CodeInvalidToken, "缺少认证头部"))
			c.Abort()
			return
		}

		// 检查Bearer token格式
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, models.Error(models.CodeInvalidToken, "无效的认证格式"))
			c.Abort()
			return
		}

		// 提取token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.JSON(http.StatusUnauthorized, models.Error(models.CodeInvalidToken, "缺少认证令牌"))
			c.Abort()
			return
		}

		// 验证token（这里简化处理，实际应该验证JWT token）
		// 在实际项目中，这里应该：
		// 1. 解析JWT token
		// 2. 验证token是否有效
		// 3. 从token中提取用户信息
		// 4. 检查用户是否存在且状态正常

		// 临时实现：假设token是用户ID
		// 实际实现应该使用JWT库来验证token
		userID := parseTokenToUserID(token)
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, models.Error(models.CodeInvalidToken, "无效的认证令牌"))
			c.Abort()
			return
		}

		// 将用户ID存储到上下文中
		c.Set("user_id", userID)
		c.Next()
	}
}

// parseTokenToUserID 解析token获取用户ID
// 这是一个临时实现，实际应该使用JWT库
func parseTokenToUserID(token string) uint {
	// 临时实现：假设token就是用户ID的字符串形式
	// 实际实现应该解析JWT token

	// 这里简化处理，返回固定的用户ID用于测试
	// 在实际项目中，应该：
	// 1. 使用JWT库解析token
	// 2. 验证token签名
	// 3. 检查token是否过期
	// 4. 从token claims中提取用户ID

	return 1 // 临时返回用户ID为1
}
