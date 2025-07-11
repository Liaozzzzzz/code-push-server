package service

import (
	"errors"
	"fmt"

	"github.com/liaozzzzzz/code-push-server/internal/dao"
	"github.com/liaozzzzzz/code-push-server/internal/models"
	"gorm.io/gorm"
)

// AppService 应用服务
type AppService struct {
	appDAO  *dao.AppDAO
	userDAO *dao.UserDAO
}

// NewAppService 创建应用服务实例
func NewAppService() *AppService {
	return &AppService{
		appDAO:  dao.NewAppDAO(),
		userDAO: dao.NewUserDAO(),
	}
}

// Create 创建应用
func (s *AppService) Create(userID uint, req *models.AppCreateRequest) (*models.AppResponse, error) {
	// 检查用户是否存在
	user, err := s.userDAO.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrUserNotFound
		}
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}

	// 检查Bundle ID是否已存在
	exists, err := s.appDAO.ExistsByBundleID(req.BundleID)
	if err != nil {
		return nil, fmt.Errorf("检查Bundle ID失败: %w", err)
	}
	if exists {
		return nil, models.ErrBundleIDExists
	}

	// 创建应用
	app := &models.App{
		Name:        req.Name,
		Platform:    req.Platform,
		BundleID:    req.BundleID,
		Description: req.Description,
		UserID:      userID,
		User:        *user,
	}

	if err := s.appDAO.Create(app); err != nil {
		return nil, fmt.Errorf("创建应用失败: %w", err)
	}

	return app.ToResponse(), nil
}

// GetByID 根据ID获取应用
func (s *AppService) GetByID(id uint) (*models.AppResponse, error) {
	app, err := s.appDAO.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrAppNotFound
		}
		return nil, fmt.Errorf("获取应用失败: %w", err)
	}

	return app.ToResponse(), nil
}

// Update 更新应用
func (s *AppService) Update(id uint, userID uint, req *models.AppUpdateRequest) (*models.AppResponse, error) {
	app, err := s.appDAO.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrAppNotFound
		}
		return nil, fmt.Errorf("获取应用失败: %w", err)
	}

	// 检查权限：只有应用所有者可以更新
	if app.UserID != userID {
		return nil, models.ErrPermissionDenied
	}

	// 更新字段
	if req.Name != "" {
		app.Name = req.Name
	}

	if req.Description != "" {
		app.Description = req.Description
	}

	if err := s.appDAO.Update(app); err != nil {
		return nil, fmt.Errorf("更新应用失败: %w", err)
	}

	return app.ToResponse(), nil
}

// Delete 删除应用
func (s *AppService) Delete(id uint, userID uint) error {
	app, err := s.appDAO.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrAppNotFound
		}
		return fmt.Errorf("获取应用失败: %w", err)
	}

	// 检查权限：只有应用所有者可以删除
	if app.UserID != userID {
		return models.ErrPermissionDenied
	}

	if err := s.appDAO.Delete(id); err != nil {
		return fmt.Errorf("删除应用失败: %w", err)
	}

	return nil
}

// List 获取应用列表
func (s *AppService) List(pageReq *models.PageRequest) ([]*models.AppResponse, int64, error) {
	apps, total, err := s.appDAO.List(pageReq.GetOffset(), pageReq.GetSize())
	if err != nil {
		return nil, 0, fmt.Errorf("获取应用列表失败: %w", err)
	}

	var responses []*models.AppResponse
	for _, app := range apps {
		responses = append(responses, app.ToResponse())
	}

	return responses, total, nil
}

// ListByUser 获取用户的应用列表
func (s *AppService) ListByUser(userID uint, pageReq *models.PageRequest) ([]*models.AppResponse, int64, error) {
	// 检查用户是否存在
	_, err := s.userDAO.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, models.ErrUserNotFound
		}
		return nil, 0, fmt.Errorf("获取用户失败: %w", err)
	}

	apps, total, err := s.appDAO.ListByUserID(userID, pageReq.GetOffset(), pageReq.GetSize())
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户应用列表失败: %w", err)
	}

	var responses []*models.AppResponse
	for _, app := range apps {
		responses = append(responses, app.ToResponse())
	}

	return responses, total, nil
}

// GetByBundleID 根据Bundle ID获取应用
func (s *AppService) GetByBundleID(bundleID string) (*models.AppResponse, error) {
	app, err := s.appDAO.GetByBundleID(bundleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrAppNotFound
		}
		return nil, fmt.Errorf("获取应用失败: %w", err)
	}

	return app.ToResponse(), nil
}
