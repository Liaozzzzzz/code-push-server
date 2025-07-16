package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/liaozzzzzz/code-push-server/internal/config"
	"github.com/liaozzzzzz/code-push-server/internal/utils/errors"
	"github.com/liaozzzzzz/code-push-server/internal/utils/response"
)

// AuthMiddleware JWT认证中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "请登录后再操作"))
			c.Abort()
			return
		}

		// 检查Bearer格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "token格式错误"))
			c.Abort()
			return
		}

		// 获取token
		token := parts[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "token不能为空"))
			c.Abort()
			return
		}

		// 验证token（这里简化处理，实际应该验证JWT）
		// 在实际项目中，这里应该：
		// 1. 解析JWT token
		// 2. 验证token的有效性
		// 3. 获取用户信息
		// 4. 设置用户上下文
		claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.C.Security.JWTSecret), nil
		})

		if err != nil || !claims.Valid {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "token无效"))
			c.Abort()
			return
		}

		// 判断token是否过期
		if exp, ok := claims.Claims.(jwt.MapClaims)["exp"]; ok {
			if time.Now().After(time.Unix(int64(exp.(float64)), 0)) {
				c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "token已过期"))
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "token已过期"))
			c.Abort()
			return
		}

		userId := claims.Claims.(jwt.MapClaims)["userId"]

		if ack, ok := claims.Claims.(jwt.MapClaims)["ack"]; ok {
			fmt.Println(ack)
			// 校验当前用户ack是否与token中的ack一致
			// user, err := service.NewUserService().GetUserByID(userId.(int))
			// if err != nil {
			// 	c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "token无效"))
			// 	c.Abort()
			// 	return
			// }
			// if ack != user.AckCode {
			// 	c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "token无效"))
			// 	c.Abort()
			// 	return
			// }
		} else {
			c.JSON(http.StatusUnauthorized, response.Error(errors.CodeInvalidToken, "token无效"))
			c.Abort()
			return
		}

		// 设置用户信息到上下文
		c.Set("userId", userId)

		c.Next()
	}
}
