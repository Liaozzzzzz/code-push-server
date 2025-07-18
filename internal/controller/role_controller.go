package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/service"
)

type RoleController struct {
	roleService *service.RoleService
}

func NewRoleController() *RoleController {
	return &RoleController{
		roleService: service.NewRoleService(),
	}
}

// @Summary      获取角色列表
// @Description  获取角色列表
// @Tags         role
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		query     dto.RoleListQueryForm  true  "role list form"
// @Success      200			{object}  response.PageResponse{data=[]dto.RoleResponse}
// @Router       /role/list		[get]
func (c *RoleController) List(ctx *gin.Context) {

}

// @Summary      创建角色
// @Description  创建角色
// @Tags         role
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.RoleCreateForm  true  "Role Create Form"
// @Success      200			{object}  response.Response{data=dto.RoleResponse}
// @Router       /role/create		[post]
func (c *RoleController) Create(ctx *gin.Context) {
}

// @Summary      更新角色
// @Description  更新角色
// @Tags         role
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.RoleUpdateForm  true  "Role Update Form"
// @Success      200			{object}  response.Response{data=nil}
// @Router       /role/update		[post]
func (c *RoleController) Update(ctx *gin.Context) {
}

// @Summary      删除角色
// @Description  删除角色
// @Tags         role
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        object		body      dto.RoleDeleteForm  true  "Role Delete Form"
// @Success      200			{object}  response.Response{data=nil}
// @Router       /role/delete		[post]
func (c *RoleController) Delete(ctx *gin.Context) {
}
