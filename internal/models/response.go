package models

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code"`    // 业务状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data,omitempty"`
}

// PageResponse 分页响应结构
type PageResponse struct {
	Code    int         `json:"code"`    // 业务状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data,omitempty"`
	Page    PageInfo    `json:"page"`
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

// 业务状态码常量
const (
	// 成功状态码
	CodeSuccess = 0

	// 客户端错误 1000-1999
	CodeInvalidParams    = 1001 // 参数错误
	CodeResourceExists   = 1002 // 资源已存在
	CodeResourceNotFound = 1003 // 资源不存在
	CodePermissionDenied = 1004 // 权限不足
	CodeLoginFailed      = 1005 // 登录失败
	CodeAccountDisabled  = 1006 // 账户被禁用
	CodeInvalidToken     = 1007 // 无效的令牌
	CodeTokenExpired     = 1008 // 令牌过期

	// 服务器错误 2000-2999
	CodeInternalError = 2001 // 内部服务器错误
	CodeDatabaseError = 2002 // 数据库错误
	CodeServiceError  = 2003 // 服务错误
)

// 业务状态码对应的消息
var CodeMessages = map[int]string{
	CodeSuccess:          "操作成功",
	CodeInvalidParams:    "参数错误",
	CodeResourceExists:   "资源已存在",
	CodeResourceNotFound: "资源不存在",
	CodePermissionDenied: "权限不足",
	CodeLoginFailed:      "登录失败",
	CodeAccountDisabled:  "账户被禁用",
	CodeInvalidToken:     "无效的令牌",
	CodeTokenExpired:     "令牌过期",
	CodeInternalError:    "内部服务器错误",
	CodeDatabaseError:    "数据库错误",
	CodeServiceError:     "服务错误",
}

// Success 成功响应
func Success(data interface{}) *Response {
	return &Response{
		Code:    CodeSuccess,
		Message: CodeMessages[CodeSuccess],
		Data:    data,
	}
}

// Error 错误响应
func Error(code int, message string) *Response {
	// 如果message为空，使用默认消息
	if message == "" {
		if defaultMsg, exists := CodeMessages[code]; exists {
			message = defaultMsg
		} else {
			message = "未知错误"
		}
	}
	return &Response{
		Code:    code,
		Message: message,
	}
}

// ErrorWithCode 使用业务状态码的错误响应
func ErrorWithCode(code int) *Response {
	return Error(code, "")
}

// PageSuccess 分页成功响应
func PageSuccess(data interface{}, page, size int, total int64) *PageResponse {
	totalPages := int(total) / size
	if int(total)%size > 0 {
		totalPages++
	}

	return &PageResponse{
		Code:    CodeSuccess,
		Message: CodeMessages[CodeSuccess],
		Data:    data,
		Page: PageInfo{
			Current:    page,
			Size:       size,
			Total:      total,
			TotalPages: totalPages,
		},
	}
}
