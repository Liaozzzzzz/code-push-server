package middleware

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// bodyWriter 用于捕获响应体
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logger 详细日志中间件，打印请求入参和返参
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()

		// 读取请求体
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				// 重新设置请求体，因为已经被读取了
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// 获取查询参数
		queryParams := c.Request.URL.RawQuery

		// 包装ResponseWriter以捕获响应体
		blw := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 计算处理时间
		latency := time.Since(start)

		// 获取响应体
		responseBody := blw.body.String()

		// 打印详细日志
		fmt.Printf(`
========== 请求详情 ==========
时间: %s
客户端IP: %s
请求方法: %s
请求路径: %s
查询参数: %s
请求头: %s
请求体: %s
状态码: %d
响应头: %s
响应体: %s
处理时间: %v
错误信息: %s
================================
`,
			start.Format("2006-01-02 15:04:05"),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			queryParams,
			formatHeaders(c.Request.Header),
			requestBody,
			c.Writer.Status(),
			formatHeaders(c.Writer.Header()),
			responseBody,
			latency,
			strings.Join(c.Errors.Errors(), "; "),
		)
	}
}

// formatHeaders 格式化请求头
func formatHeaders(headers map[string][]string) string {
	var headerStrings []string
	for key, values := range headers {
		headerStrings = append(headerStrings, fmt.Sprintf("%s: %s", key, strings.Join(values, ", ")))
	}
	return strings.Join(headerStrings, "; ")
}
