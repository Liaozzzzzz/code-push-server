package service

import (
	"errors"
	"fmt"

	"github.com/liaozzzzzz/code-push-server/internal/dao"
	"github.com/liaozzzzzz/code-push-server/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 用户服务
type UserService struct {
	userDAO *dao.UserDAO
}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{
		userDAO: dao.NewUserDAO(),
	}
}

// Create 创建用户
func (s *UserService) Create(req *models.UserCreateRequest) (*models.UserResponse, error) {
	// 检查用户名是否已存在
	exists, err := s.userDAO.ExistsByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("检查用户名失败: %w", err)
	}
	if exists {
		return nil, models.ErrUsernameExists
	}

	// 检查邮箱是否已存在
	exists, err = s.userDAO.ExistsByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("检查邮箱失败: %w", err)
	}
	if exists {
		return nil, models.ErrEmailExists
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %w", err)
	}

	// 设置默认角色
	role := req.Role
	if role == "" {
		role = "developer"
	}

	// 创建用户
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
		Status:   "active",
	}

	if err := s.userDAO.Create(user); err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	return user.ToResponse(), nil
}

// GetByID 根据ID获取用户
func (s *UserService) GetByID(id uint) (*models.UserResponse, error) {
	user, err := s.userDAO.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrUserNotFound
		}
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}

	return user.ToResponse(), nil
}

// Update 更新用户
func (s *UserService) Update(id uint, req *models.UserUpdateRequest) (*models.UserResponse, error) {
	user, err := s.userDAO.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrUserNotFound
		}
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}

	// 更新字段
	if req.Username != "" {
		// 检查用户名是否已被其他用户使用
		existingUser, err := s.userDAO.GetByUsername(req.Username)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("检查用户名失败: %w", err)
		}
		if existingUser != nil && existingUser.ID != id {
			return nil, models.ErrUsernameExists
		}
		user.Username = req.Username
	}

	if req.Email != "" {
		// 检查邮箱是否已被其他用户使用
		existingUser, err := s.userDAO.GetByEmail(req.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("检查邮箱失败: %w", err)
		}
		if existingUser != nil && existingUser.ID != id {
			return nil, models.ErrEmailExists
		}
		user.Email = req.Email
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	if req.Status != "" {
		user.Status = req.Status
	}

	if err := s.userDAO.Update(user); err != nil {
		return nil, fmt.Errorf("更新用户失败: %w", err)
	}

	return user.ToResponse(), nil
}

// Delete 删除用户
func (s *UserService) Delete(id uint) error {
	// 检查用户是否存在
	_, err := s.userDAO.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrUserNotFound
		}
		return fmt.Errorf("获取用户失败: %w", err)
	}

	if err := s.userDAO.Delete(id); err != nil {
		return fmt.Errorf("删除用户失败: %w", err)
	}

	return nil
}

// List 获取用户列表
func (s *UserService) List(pageReq *models.PageRequest) ([]*models.UserResponse, int64, error) {
	users, total, err := s.userDAO.List(pageReq.GetOffset(), pageReq.GetSize())
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户列表失败: %w", err)
	}

	var responses []*models.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	return responses, total, nil
}

// Login 用户登录
func (s *UserService) Login(username, password string) (*models.UserResponse, error) {
	user, err := s.userDAO.GetByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrLoginFailed
		}
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}

	// 检查用户状态
	if user.Status != "active" {
		return nil, models.ErrAccountDisabled
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, models.ErrLoginFailed
	}

	return user.ToResponse(), nil
}

// UpdateStatus 更新用户状态
func (s *UserService) UpdateStatus(id uint, status string) error {
	// 检查用户是否存在
	_, err := s.userDAO.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrUserNotFound
		}
		return fmt.Errorf("获取用户失败: %w", err)
	}

	if err := s.userDAO.UpdateStatus(id, status); err != nil {
		return fmt.Errorf("更新用户状态失败: %w", err)
	}

	return nil
}
