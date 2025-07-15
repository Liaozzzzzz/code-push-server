package response

import (
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
