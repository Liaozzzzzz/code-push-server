package dto

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/liaozzzzzz/code-push-server/internal/entity"
	"github.com/liaozzzzzz/code-push-server/internal/types"
)

// DeptCreateForm 创建部门表单
type DeptCreateForm struct {
	DeptName string           `json:"deptName" binding:"required,min=3,max=50"`
	ParentID *int64           `json:"parentId" binding:"required"`
	Sort     *int32           `json:"sort" binding:"required"`
	Leader   *string          `json:"leader" binding:"omitempty,max=50"`
	Phone    *string          `json:"phone" binding:"omitempty,max=50"`
	Email    *string          `json:"email" binding:"omitempty,max=50"`
	Status   types.DeptStatus `json:"status" binding:"required,oneof='1' '0'"`
}

type DeptUpdateForm struct {
	DeptID int64 `json:"deptId" binding:"required"`
	DeptCreateForm
}

type DeptDeleteForm struct {
	DeptID int64 `json:"deptId" binding:"required"`
}

// DeptResponse 部门响应
type DeptResponse struct {
	DeptID    int64            `json:"deptId"`
	ParentID  int64            `json:"parentId"`
	DeptName  string           `json:"deptName"`
	Sort      int32            `json:"sort"`
	Leader    *string          `json:"leader"`
	Phone     *string          `json:"phone"`
	Email     *string          `json:"email"`
	Status    types.DeptStatus `json:"status"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

// DeptTreeResponse 部门树形结构响应
type DeptTreeResponse struct {
	DeptResponse
	Children []*DeptTreeResponse `json:"children"`
}

func BuildDeptTree(deptList []*entity.Dept, parentID int64) []*DeptTreeResponse {
	deptTree := make([]*DeptTreeResponse, 0)
	for _, dept := range deptList {
		if dept.ParentID != parentID {
			continue
		}
		var deptResponse DeptTreeResponse
		if err := copier.Copy(&deptResponse, dept); err != nil {
			continue
		}
		deptResponse.Children = BuildDeptTree(deptList, dept.DeptID)
		deptTree = append(deptTree, &deptResponse)
	}
	return deptTree
}
