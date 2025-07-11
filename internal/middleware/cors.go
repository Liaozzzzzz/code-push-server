package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/config"
)

// CORS 跨域中间件
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		corsConfig := config.C.Security.CORS

		// 设置允许的来源
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 简单实现：如果配置了*，则允许所有来源
			if len(corsConfig.AllowedOrigins) > 0 && corsConfig.AllowedOrigins[0] == "*" {
				c.Header("Access-Control-Allow-Origin", "*")
			} else {
				// 检查是否在允许列表中
				for _, allowedOrigin := range corsConfig.AllowedOrigins {
					if origin == allowedOrigin {
						c.Header("Access-Control-Allow-Origin", origin)
						break
					}
				}
			}
		}

		// 设置允许的方法
		if len(corsConfig.AllowedMethods) > 0 {
			methods := ""
			for i, method := range corsConfig.AllowedMethods {
				if i > 0 {
					methods += ", "
				}
				methods += method
			}
			c.Header("Access-Control-Allow-Methods", methods)
		}

		// 设置允许的头部
		if len(corsConfig.AllowedHeaders) > 0 {
			headers := ""
			for i, header := range corsConfig.AllowedHeaders {
				if i > 0 {
					headers += ", "
				}
				headers += header
			}
			c.Header("Access-Control-Allow-Headers", headers)
		}

		// 设置暴露的头部
		if len(corsConfig.ExposedHeaders) > 0 {
			headers := ""
			for i, header := range corsConfig.ExposedHeaders {
				if i > 0 {
					headers += ", "
				}
				headers += header
			}
			c.Header("Access-Control-Expose-Headers", headers)
		}

		// 设置是否允许凭证
		if corsConfig.AllowCredentials {
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// 设置预检请求的缓存时间
		if corsConfig.MaxAge > 0 {
			c.Header("Access-Control-Max-Age", string(rune(corsConfig.MaxAge)))
		}

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
