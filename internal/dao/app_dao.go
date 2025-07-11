package dao

import (
	"github.com/liaozzzzzz/code-push-server/internal/database"
	"github.com/liaozzzzzz/code-push-server/internal/models"
	"gorm.io/gorm"
)

// AppDAO 应用数据访问对象
type AppDAO struct {
	db *gorm.DB
}

// NewAppDAO 创建应用DAO实例
func NewAppDAO() *AppDAO {
	return &AppDAO{
		db: database.DB,
	}
}

// Create 创建应用
func (d *AppDAO) Create(app *models.App) error {
	return d.db.Create(app).Error
}

// GetByID 根据ID获取应用
func (d *AppDAO) GetByID(id uint) (*models.App, error) {
	var app models.App
	err := d.db.Preload("User").First(&app, id).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

// GetByBundleID 根据Bundle ID获取应用
func (d *AppDAO) GetByBundleID(bundleID string) (*models.App, error) {
	var app models.App
	err := d.db.Preload("User").Where("bundle_id = ?", bundleID).First(&app).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

// Update 更新应用
func (d *AppDAO) Update(app *models.App) error {
	return d.db.Save(app).Error
}

// Delete 删除应用
func (d *AppDAO) Delete(id uint) error {
	return d.db.Delete(&models.App{}, id).Error
}

// List 获取应用列表
func (d *AppDAO) List(offset, limit int) ([]*models.App, int64, error) {
	var apps []*models.App
	var total int64

	// 获取总数
	if err := d.db.Model(&models.App{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := d.db.Preload("User").Offset(offset).Limit(limit).Find(&apps).Error
	if err != nil {
		return nil, 0, err
	}

	return apps, total, nil
}

// ListByUserID 根据用户ID获取应用列表
func (d *AppDAO) ListByUserID(userID uint, offset, limit int) ([]*models.App, int64, error) {
	var apps []*models.App
	var total int64

	// 获取总数
	if err := d.db.Model(&models.App{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := d.db.Preload("User").Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&apps).Error
	if err != nil {
		return nil, 0, err
	}

	return apps, total, nil
}

// ExistsByBundleID 检查Bundle ID是否存在
func (d *AppDAO) ExistsByBundleID(bundleID string) (bool, error) {
	var count int64
	err := d.db.Model(&models.App{}).Where("bundle_id = ?", bundleID).Count(&count).Error
	return count > 0, err
}

// ExistsByUserAndBundleID 检查用户是否拥有指定Bundle ID的应用
func (d *AppDAO) ExistsByUserAndBundleID(userID uint, bundleID string) (bool, error) {
	var count int64
	err := d.db.Model(&models.App{}).Where("user_id = ? AND bundle_id = ?", userID, bundleID).Count(&count).Error
	return count > 0, err
}
