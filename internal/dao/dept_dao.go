package dao

import (
	"github.com/liaozzzzzz/code-push-server/internal/entity"
	"github.com/liaozzzzzz/code-push-server/internal/utils/database"
	utilsErrors "github.com/liaozzzzzz/code-push-server/internal/utils/errors"
	"gorm.io/gorm"
)

type DeptDAO struct {
	db *gorm.DB
}

func NewDeptDAO() *DeptDAO {
	return &DeptDAO{
		db: database.DB,
	}
}

// 查询部门列表
func (d *DeptDAO) GetDeptList() ([]*entity.Dept, error) {
	var depts []*entity.Dept
	result := d.db.Find(&depts)
	if result.Error != nil {
		return nil, result.Error
	}

	return depts, nil
}

// 根据id查询部门
func (d *DeptDAO) GetDeptByID(id int64) (*entity.Dept, error) {
	var dept entity.Dept
	result := d.db.First(&dept, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dept, nil
}

// 新增部门
func (d *DeptDAO) Create(dept *entity.Dept) (*entity.Dept, error) {
	result := d.db.Create(&dept)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, utilsErrors.NewBusinessErrorf(utilsErrors.CodeCreateFailed, "dept create failed")
	}
	return dept, nil
}

func (d *DeptDAO) Update(dept *entity.Dept) error {
	// 使用 Updates 方法，只更新非零值字段，避免时间字段的零值问题
	result := d.db.Model(dept).Updates(dept)
	if result.RowsAffected == 0 {
		return utilsErrors.NewBusinessErrorf(utilsErrors.CodeUpdateFailed, "dept update failed")
	}
	return result.Error
}

// 批量删除部门
func (d *DeptDAO) BatchDelete(deptIDs []int64) error {
	return d.db.Where("dept_id IN (?)", deptIDs).Delete(&entity.Dept{}).Error
}
