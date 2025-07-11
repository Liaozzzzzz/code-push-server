package models

import "fmt"

// BusinessError 业务错误
type BusinessError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error 实现error接口
func (e *BusinessError) Error() string {
	return e.Message
}

// GetCode 获取业务状态码
func (e *BusinessError) GetCode() int {
	return e.Code
}

// NewBusinessError 创建业务错误
func NewBusinessError(code int, message string) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
	}
}

// 预定义的业务错误
var (
	// 资源不存在错误
	ErrUserNotFound = NewBusinessError(CodeResourceNotFound, "用户不存在")
	ErrAppNotFound  = NewBusinessError(CodeResourceNotFound, "应用不存在")

	// 资源已存在错误
	ErrUsernameExists = NewBusinessError(CodeResourceExists, "用户名已存在")
	ErrEmailExists    = NewBusinessError(CodeResourceExists, "邮箱已存在")
	ErrBundleIDExists = NewBusinessError(CodeResourceExists, "Bundle ID已存在")

	// 权限错误
	ErrPermissionDenied = NewBusinessError(CodePermissionDenied, "无权限操作此应用")

	// 认证错误
	ErrLoginFailed     = NewBusinessError(CodeLoginFailed, "用户名或密码错误")
	ErrAccountDisabled = NewBusinessError(CodeAccountDisabled, "用户账户已被禁用")

	// 服务错误
	ErrServiceError = NewBusinessError(CodeServiceError, "服务错误")
)

// NewBusinessErrorf 创建带格式化消息的业务错误
func NewBusinessErrorf(code int, format string, args ...interface{}) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
