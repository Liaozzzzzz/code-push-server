package dao

import (
	"github.com/liaozzzzzz/code-push-server/internal/entity"
	"github.com/liaozzzzzz/code-push-server/internal/utils/database"
	utilsErrors "github.com/liaozzzzzz/code-push-server/internal/utils/errors"
	"gorm.io/gorm"
)

type MenuDAO struct {
	db *gorm.DB
}

func NewMenuDAO() *MenuDAO {
	return &MenuDAO{
		db: database.DB,
	}
}

func (d *MenuDAO) GetMenuList() ([]*entity.Menu, error) {
	var menuList []*entity.Menu
	if err := d.db.Find(&menuList).Error; err != nil {
		return nil, err
	}
	return menuList, nil
}

// Create 创建菜单
func (d *MenuDAO) Create(menu *entity.Menu) (*entity.Menu, error) {
	if err := d.db.Create(menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}

func (d *MenuDAO) GetByID(id int64) (*entity.Menu, error) {
	var menu entity.Menu
	if err := d.db.First(&menu, id).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (d *MenuDAO) Update(menu *entity.Menu) error {
	result := d.db.Model(menu).Where("menu_id = ?", menu.MenuID).Select("*").Omit("created_at", "updated_at", "deleted_at").Updates(menu)
	if result.RowsAffected == 0 {
		return utilsErrors.NewBusinessErrorf(utilsErrors.CodeUpdateFailed, "menu update failed")
	}
	return result.Error
}

func (d *MenuDAO) Delete(id int64) error {
	return d.db.Where("menu_id = ?", id).Delete(&entity.Menu{}).Error
}
