package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/liaozzzzzz/code-push-server/internal/dao"
	"github.com/liaozzzzzz/code-push-server/internal/dto"
	"github.com/liaozzzzzz/code-push-server/internal/entity"
	utilsErrors "github.com/liaozzzzzz/code-push-server/internal/utils/errors"
)

type MenuService struct {
	menuRepository *dao.MenuDAO
}

func NewMenuService() *MenuService {
	return &MenuService{
		menuRepository: dao.NewMenuDAO(),
	}
}

func (s *MenuService) Tree() ([]*dto.MenuTreeResponse, error) {
	menuList, err := s.menuRepository.GetMenuList()
	if err != nil {
		return nil, err
	}
	return dto.BuildMenuTree(menuList, 0), nil
}

func (s *MenuService) Create(ctx context.Context, form dto.MenuCreateForm) (*dto.MenuResponse, error) {
	// 父菜单是否存在
	if *form.ParentID != 0 {
		if parentMenu, err := s.menuRepository.GetByID(*form.ParentID); err != nil {
			return nil, err
		} else if parentMenu == nil {
			return nil, utilsErrors.NewBusinessErrorf(utilsErrors.CodeResourceNotFound, "parent menu not found")
		}
	}

	var menu entity.Menu
	if err := copier.Copy(&menu, form); err != nil {
		return nil, err
	}

	createdMenu, err := s.menuRepository.Create(&menu)
	if err != nil {
		return nil, err
	}
	var response dto.MenuResponse
	if err := copier.Copy(&response, createdMenu); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *MenuService) Update(ctx context.Context, form dto.MenuUpdateForm) error {
	// 菜单是否存在
	if menu, err := s.menuRepository.GetByID(form.MenuID); err != nil {
		return err
	} else if menu == nil {
		return utilsErrors.NewBusinessErrorf(utilsErrors.CodeResourceNotFound, "menu not found")
	}

	// 父菜单是否存在
	if *form.ParentID != 0 {
		if parentMenu, err := s.menuRepository.GetByID(*form.ParentID); err != nil {
			return err
		} else if parentMenu == nil {
			return utilsErrors.NewBusinessErrorf(utilsErrors.CodeResourceNotFound, "parent menu not found")
		}
	}

	var menu entity.Menu
	if err := copier.Copy(&menu, form); err != nil {
		return err
	}

	return s.menuRepository.Update(&menu)
}

func (s *MenuService) Delete(ctx context.Context, form dto.MenuDeleteForm) error {
	// 菜单是否存在
	if menu, err := s.menuRepository.GetByID(form.MenuID); err != nil {
		return err
	} else if menu == nil {
		return utilsErrors.NewBusinessErrorf(utilsErrors.CodeResourceNotFound, "menu not found")
	}
	return s.menuRepository.Delete(form.MenuID)
}
