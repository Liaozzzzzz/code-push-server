package types

// Status 通用状态类型
type Status string

const (
	// StatusEnabled 启用状态
	StatusEnabled Status = "1"
	// StatusDisabled 禁用状态
	StatusDisabled Status = "0"
)

// UserStatus 用户状态类型
type UserStatus Status

const (
	// UserEnabled 用户启用
	UserEnabled UserStatus = UserStatus(StatusEnabled)
	// UserDisabled 用户禁用
	UserDisabled UserStatus = UserStatus(StatusDisabled)
)

// RoleStatus 角色状态类型
type RoleStatus Status

const (
	// RoleEnabled 角色启用
	RoleEnabled RoleStatus = RoleStatus(StatusEnabled)
	// RoleDisabled 角色禁用
	RoleDisabled RoleStatus = RoleStatus(StatusDisabled)
)

// MenuStatus 菜单状态类型
type MenuStatus Status

const (
	// MenuEnabled 菜单启用
	MenuEnabled MenuStatus = MenuStatus(StatusEnabled)
	// MenuDisabled 菜单禁用
	MenuDisabled MenuStatus = MenuStatus(StatusDisabled)
)

// MenuType 菜单类型
type MenuType string

const (
	// MenuTypeDirectory 目录
	MenuTypeDirectory MenuType = "1"
	// MenuTypeMenu 菜单
	MenuTypeMenu MenuType = "2"
	// MenuTypeButton 按钮
	MenuTypeButton MenuType = "3"
)

// MenuVisible 菜单显示状态
type MenuVisible Status

const (
	// MenuVisible 显示
	MenuVisibleShow MenuVisible = MenuVisible(StatusEnabled)
	// MenuHidden 隐藏
	MenuVisibleHidden MenuVisible = MenuVisible(StatusDisabled)
)

// MenuIsLink 是否外链
type MenuIsLink Status

const (
	// MenuIsExternalLink 外链
	MenuIsExternalLink MenuIsLink = MenuIsLink(StatusEnabled)
	// MenuIsInternalLink 内链
	MenuIsInternalLink MenuIsLink = MenuIsLink(StatusDisabled)
)
