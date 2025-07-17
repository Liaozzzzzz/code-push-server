package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/dto"
	"github.com/liaozzzzzz/code-push-server/internal/service"
	"github.com/liaozzzzzz/code-push-server/internal/utils/response"
)

type MenuController struct {
	menuService *service.MenuService
}

func NewMenuController() *MenuController {
	return &MenuController{
		menuService: service.NewMenuService(),
	}
}

// @Summary      获取菜单树
// @Description  获取菜单树
// @Tags         menu
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200		{object}  response.Response{data=[]dto.MenuTreeResponse}
// @Router       /menu/tree		[get]
func (c *MenuController) Tree(ctx *gin.Context) {
	tree, err := c.menuService.Tree()
	if err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, tree)
}

// @Summary      创建菜单
// @Description  创建菜单
// @Tags         menu
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.MenuCreateForm  true  "Menu Create Form"
// @Success      200		{object}  response.Response{data=dto.MenuResponse}
// @Router       /menu/create		[post]
func (c *MenuController) Create(ctx *gin.Context) {
	var form dto.MenuCreateForm
	if err := response.ParseJSON(ctx, &form); err != nil {
		response.HandleError(ctx, err)
		return
	}

	menu, err := c.menuService.Create(ctx, form)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, menu)
}

// @Summary      更新菜单
// @Description  更新菜单
// @Tags         menu
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.MenuUpdateForm  true  "Menu Update Form"
// @Success      200		{object}  response.Response{data=nil}
// @Router       /menu/update		[put]
func (c *MenuController) Update(ctx *gin.Context) {
	var form dto.MenuUpdateForm
	if err := response.ParseJSON(ctx, &form); err != nil {
		response.HandleError(ctx, err)
		return
	}
	if err := c.menuService.Update(ctx, form); err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, nil)
}

// @Summary      删除菜单
// @Description  删除菜单
// @Tags         menu
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.MenuDeleteForm  true  "Menu Delete Form"
// @Success      200		{object}  response.Response{data=nil}
// @Router       /menu/delete		[delete]
func (c *MenuController) Delete(ctx *gin.Context) {
	var form dto.MenuDeleteForm
	if err := response.ParseJSON(ctx, &form); err != nil {
		response.HandleError(ctx, err)
		return
	}
	if err := c.menuService.Delete(ctx, form); err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, nil)
}
