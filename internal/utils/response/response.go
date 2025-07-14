package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/liaozzzzzz/code-push-server/internal/utils/errors"
)

// Response 通用响应结构
type Response struct {
	Code    errors.BusinessCode `json:"code"`    // 业务状态码
	Message string              `json:"message"` // 响应消息
	Data    interface{}         `json:"data,omitempty"`
}

// PageResponse 分页响应结构
type PageResponse struct {
	Code    errors.BusinessCode `json:"code"`    // 业务状态码
	Message string              `json:"message"` // 响应消息
	Data    interface{}         `json:"data,omitempty"`
	Page    PageInfo            `json:"page"`
}

// PageInfo 分页信息
type PageInfo struct {
	Current    int   `json:"current"`
	Size       int   `json:"size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// PageRequest 分页请求
type PageRequest struct {
	Page int `form:"page" binding:"omitempty,min=1"`
	Size int `form:"size" binding:"omitempty,min=1,max=100"`
}

// GetPage 获取页码，默认为1
func (p *PageRequest) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

// GetSize 获取每页大小，默认为10
func (p *PageRequest) GetSize() int {
	if p.Size <= 0 {
		return 10
	}
	return p.Size
}

// GetOffset 获取偏移量
func (p *PageRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

// Success 成功响应
func Success(data interface{}) *Response {
	return &Response{
		Code:    errors.CodeSuccess,
		Message: "操作成功",
		Data:    data,
	}
}

// Error 错误响应
func Error(code errors.BusinessCode, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

// FromBusinessError 从BusinessError创建Response
func FromBusinessError(err *errors.BusinessError) *Response {
	return &Response{
		Code:    err.GetCode(),
		Message: err.GetMessage(),
	}
}

// PageSuccess 分页成功响应
func PageSuccess(data interface{}, page, size int, total int64) *PageResponse {
	totalPages := int(total) / size
	if int(total)%size > 0 {
		totalPages++
	}

	return &PageResponse{
		Code:    errors.CodeSuccess,
		Message: "操作成功",
		Data:    data,
		Page: PageInfo{
			Current:    page,
			Size:       size,
			Total:      total,
			TotalPages: totalPages,
		},
	}
}

// HandleParamError 处理参数错误
func HandleParamError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Error(errors.CodeInvalidParams, message))
}

// HandleSuccess 处理成功响应
func HandleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Success(data))
}

// HandlePageSuccess 处理分页成功响应
func HandlePageSuccess(c *gin.Context, data interface{}, page, size int, total int64) {
	c.JSON(http.StatusOK, PageSuccess(data, page, size, total))
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

// HandleBusinessError 处理业务错误（返回200状态码）
func HandleBusinessError(c *gin.Context, code errors.BusinessCode, message string) {
	c.JSON(http.StatusOK, Error(code, message))
}

// HandleBusinessErrorFromError 处理BusinessError类型的错误
func HandleBusinessErrorFromError(c *gin.Context, err *errors.BusinessError) {
	c.JSON(http.StatusOK, FromBusinessError(err))
}

// Parse body json data to struct
func ParseJSON(c *gin.Context, obj interface{}) *errors.BusinessError {
	if err := c.ShouldBindJSON(obj); err != nil {
		fmt.Println(err)
		return errors.NewBusinessErrorf(errors.CodeInvalidParams, "Failed to parse json: %s", err.Error())
	}
	return nil
}

// Parse query parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) *errors.BusinessError {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.NewBusinessErrorf(errors.CodeInvalidParams, "Failed to parse query: %s", err.Error())
	}
	return nil
}

// Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) *errors.BusinessError {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.NewBusinessErrorf(errors.CodeInvalidParams, "Failed to parse form: %s", err.Error())
	}
	return nil
}
