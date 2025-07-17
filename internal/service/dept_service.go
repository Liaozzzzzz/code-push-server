package service

import (
	"github.com/jinzhu/copier"
	"github.com/liaozzzzzz/code-push-server/internal/dao"
	"github.com/liaozzzzzz/code-push-server/internal/dto"
	"github.com/liaozzzzzz/code-push-server/internal/entity"
	"github.com/liaozzzzzz/code-push-server/internal/types"
	utilsErrors "github.com/liaozzzzzz/code-push-server/internal/utils/errors"
)

type DeptService struct {
	deptRepository *dao.DeptDAO
}

func NewDeptService() *DeptService {
	return &DeptService{
		deptRepository: dao.NewDeptDAO(),
	}
}

// 递归获取所有子部门
func getChildrenDeptList(deptList []*entity.Dept, parentID int64) []int64 {
	if parentID == 0 {
		return []int64{}
	}

	childrenDeptIDs := []int64{parentID}
	for _, dept := range deptList {
		if dept.ParentID != parentID {
			continue
		}
		childrenDeptIDs = append(childrenDeptIDs, getChildrenDeptList(deptList, dept.DeptID)...)
	}

	return childrenDeptIDs
}

func (s *DeptService) SelectDeptTree() ([]*dto.DeptTreeResponse, error) {
	deptList, err := s.deptRepository.GetDeptList()
	if err != nil {
		return nil, err
	}

	deptTree := dto.BuildDeptTree(deptList, 0)
	return deptTree, nil
}

func (s *DeptService) Create(form dto.DeptCreateForm) (*dto.DeptResponse, error) {
	// 判断部门是否存在
	parentDept, err := s.deptRepository.GetDeptByID(*form.ParentID)
	if err != nil {
		return nil, err
	}
	if parentDept == nil {
		return nil, utilsErrors.NewBusinessError(utilsErrors.CodeResourceNotFound, "父部门不存在")
	}
	if parentDept.Status == types.DeptDisabled {
		return nil, utilsErrors.NewBusinessError(utilsErrors.CodeDisabled, "父部门已禁用")
	}

	var dept entity.Dept
	if err := copier.Copy(&dept, form); err != nil {
		return nil, err
	}

	createdDept, err := s.deptRepository.Create(&dept)
	if err != nil {
		return nil, err
	}

	var response dto.DeptResponse
	if err := copier.Copy(&response, createdDept); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *DeptService) Update(form dto.DeptUpdateForm) error {
	var dept entity.Dept
	if err := copier.Copy(&dept, form); err != nil {
		return err
	}

	return s.deptRepository.Update(&dept)
}

func (s *DeptService) Delete(form dto.DeptDeleteForm) error {
	// 删除所有子部门
	deptList, err := s.deptRepository.GetDeptList()
	if err != nil {
		return err
	}

	childrenDeptIDs := getChildrenDeptList(deptList, form.DeptID)

	return s.deptRepository.BatchDelete(childrenDeptIDs)
}
