package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/dto"
	"github.com/liaozzzzzz/code-push-server/internal/service"
	"github.com/liaozzzzzz/code-push-server/internal/utils/response"
)

type DeptController struct {
	deptService *service.DeptService
}

func NewDeptController() *DeptController {
	return &DeptController{
		deptService: service.NewDeptService(),
	}
}

// @Summary      获取部门树
// @Description  获取部门树形结构
// @Tags         dept
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200			{object}  response.Response{data=[]dto.DeptTreeResponse}
// @Router       /dept/tree		[get]
func (c *DeptController) SelectDeptTree(ctx *gin.Context) {
	deptTree, err := c.deptService.SelectDeptTree()
	if err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, deptTree)
}

// @Summary      创建部门
// @Description  创建部门
// @Tags         dept
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.DeptCreateForm  true  "Dept Create Form"
// @Success      200		{object}  response.Response{data=dto.DeptResponse}
// @Router       /dept/create		[post]
func (c *DeptController) Create(ctx *gin.Context) {
	var form dto.DeptCreateForm
	if err := response.ParseJSON(ctx, &form); err != nil {
		response.HandleError(ctx, err)
		return
	}

	deptResponse, err := c.deptService.Create(form)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, deptResponse)
}

// @Summary      更新部门信息
// @Description  更新部门信息
// @Tags         dept
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.DeptUpdateForm  true  "Dept Update Form"
// @Success      200		{object}  response.Response{data=nil}
// @Router       /dept/update		[put]
func (c *DeptController) Update(ctx *gin.Context) {
	var form dto.DeptUpdateForm
	if err := response.ParseJSON(ctx, &form); err != nil {
		response.HandleError(ctx, err)
		return
	}

	err := c.deptService.Update(form)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, nil)
}

// @Summary      删除部门
// @Description  删除部门
// @Tags         dept
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.DeptDeleteForm  true  "Dept Delete Form"
// @Success      200		{object}  response.Response{data=nil}
// @Router       /dept/delete		[delete]
func (c *DeptController) Delete(ctx *gin.Context) {
	var form dto.DeptDeleteForm
	if err := response.ParseJSON(ctx, &form); err != nil {
		response.HandleError(ctx, err)
		return
	}

	err := c.deptService.Delete(form)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, nil)
}
