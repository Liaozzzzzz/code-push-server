package entity

import (
	"time"

	"github.com/liaozzzzzz/code-push-server/internal/types"
	"gorm.io/gorm"
)

type Dept struct {
	DeptID    int64            `json:"deptId" gorm:"type:bigint(20);primaryKey;autoIncrement;comment:'部门ID'"`
	ParentID  int64            `json:"parentId" gorm:"type:bigint(20);default:0;comment:'父部门ID'"`
	DeptName  string           `json:"deptName" gorm:"type:varchar(50);not null;uniqueIndex;comment:'部门名称'"`
	Sort      *int32           `json:"sort" gorm:"type:int(11);not null;default:0;comment:'部门排序'"`
	Leader    *string          `json:"leader" gorm:"type:varchar(50);default:null;comment:'部门领导'"`
	Phone     *string          `json:"phone" gorm:"type:varchar(50);default:null;comment:'部门电话'"`
	Email     *string          `json:"email" gorm:"type:varchar(50);default:null;comment:'部门邮箱'"`
	Status    types.DeptStatus `json:"status" gorm:"type:enum('1','0');default:'1';comment:'状态'"`
	CreatedAt time.Time        `json:"createdAt" gorm:"comment:'创建时间'"`
	UpdatedAt time.Time        `json:"updatedAt" gorm:"comment:'更新时间'"`
	DeletedAt gorm.DeletedAt   `json:"_" gorm:"comment:'删除时间'"`
}

func (Dept) TableName() string {
	return "depts"
}
