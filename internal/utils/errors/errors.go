package errors

import (
	"fmt"
)

// BusinessError 业务错误
type BusinessError struct {
	Code    BusinessCode `json:"code"`
	Message string       `json:"message"`
}

// Error 实现error接口
func (e *BusinessError) Error() string {
	return e.Message
}

// GetCode 获取业务状态码
func (e *BusinessError) GetCode() BusinessCode {
	return e.Code
}

// GetMessage 获取错误消息
func (e *BusinessError) GetMessage() string {
	return e.Message
}

// NewBusinessError 创建业务错误
// 如果不传message或message为空，使用默认消息
func NewBusinessError(code BusinessCode, message ...string) *BusinessError {
	var msg string
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	} else {
		msg = code.GetMessage()
	}
	return &BusinessError{
		Code:    code,
		Message: msg,
	}
}

// NewBusinessErrorf 创建带格式化消息的业务错误
func NewBusinessErrorf(code BusinessCode, format string, args ...interface{}) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
