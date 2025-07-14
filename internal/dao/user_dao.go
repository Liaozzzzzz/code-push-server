package dao

import (
	"github.com/liaozzzzzz/code-push-server/internal/entity"
	"github.com/liaozzzzzz/code-push-server/internal/utils/database"
	"gorm.io/gorm"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 创建用户DAO实例
func NewUserDAO() *UserDAO {
	return &UserDAO{
		db: database.DB,
	}
}

// Create 创建用户
func (d *UserDAO) Create(user *entity.User) error {
	return d.db.Create(user).Error
}

// GetByID 根据ID获取用户
func (d *UserDAO) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	err := d.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
func (d *UserDAO) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱获取用户
func (d *UserDAO) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := d.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
func (d *UserDAO) Update(user *entity.User) error {
	return d.db.Save(user).Error
}

// Delete 删除用户
func (d *UserDAO) Delete(id uint) error {
	return d.db.Delete(&entity.User{}, id).Error
}

// List 获取用户列表
func (d *UserDAO) List(offset, limit int) ([]*entity.User, int64, error) {
	var users []*entity.User
	var total int64

	// 获取总数
	if err := d.db.Model(&entity.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := d.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// ExistsByUsername 检查用户名是否存在
func (d *UserDAO) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := d.db.Model(&entity.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// ExistsByEmail 检查邮箱是否存在
func (d *UserDAO) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := d.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// UpdateStatus 更新用户状态
func (d *UserDAO) UpdateStatus(id uint, status entity.UserStatus) error {
	return d.db.Model(&entity.User{}).Where("id = ?", id).Update("status", status).Error
}
