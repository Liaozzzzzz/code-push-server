package errors

// BusinessCode 业务状态码
type BusinessCode int

// 业务状态码常量
const (
	// 成功状态码
	CodeSuccess BusinessCode = 10000

	// 业务错误 20000-29999
	CodeInvalidParams    BusinessCode = 20001 // 参数错误
	CodeResourceExists   BusinessCode = 20002 // 资源已存在
	CodeResourceNotFound BusinessCode = 20003 // 资源不存在
	CodePermissionDenied BusinessCode = 20004 // 权限不足
	CodeInvalidToken     BusinessCode = 20007 // 无效的令牌
	CodeTokenExpired     BusinessCode = 20008 // 令牌过期
	CodeCreateFailed     BusinessCode = 20009 // 创建失败
	CodeUpdateFailed     BusinessCode = 20010 // 更新失败
	CodeDisabled         BusinessCode = 20011 // 已禁用

	// 服务器错误 30000-39999
	CodeInternalError BusinessCode = 30001 // 内部服务器错误
	CodeDatabaseError BusinessCode = 30002 // 数据库错误
	CodeServiceError  BusinessCode = 30003 // 服务错误
)

// 业务状态码对应的消息
var CodeMessages = map[BusinessCode]string{
	CodeSuccess:          "操作成功",
	CodeInvalidParams:    "参数错误",
	CodeResourceExists:   "资源已存在",
	CodeResourceNotFound: "资源不存在",
	CodePermissionDenied: "权限不足",
	CodeInvalidToken:     "无效的令牌",
	CodeTokenExpired:     "令牌过期",
	CodeInternalError:    "内部服务器错误",
	CodeDatabaseError:    "数据库错误",
	CodeServiceError:     "服务错误",
}

// GetMessage 获取状态码对应的消息
func (c BusinessCode) GetMessage() string {
	if msg, exists := CodeMessages[c]; exists {
		return msg
	}
	return "未知错误"
}
