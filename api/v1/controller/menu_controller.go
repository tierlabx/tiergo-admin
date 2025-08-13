package controller

import (
	"net/http"
	"tier-up/internal/app/service"

	"github.com/gin-gonic/gin"
)

// RoleController 菜单控制器
type MenuController struct {
	MenuService *service.MenuService
}

// NewRoleController 创建菜单控制器
func NewMenuController(menuService *service.MenuService) *MenuController {
	return &MenuController{
		MenuService: menuService,
	}
}

// GetRoleByID 获取菜单
// @Summary 获取菜单树
// @Tags Menu
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Menu "菜单详情"
// @Router /menu/tree [get]
func (m *MenuController) GetMenuTree(ctx *gin.Context) {
	tree, err := m.MenuService.Tree()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": tree})
}

// GetUserPermissionMenuTree 获取用户权限菜单树
// @Summary 获取用户权限菜单树
// @Tags Menu
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Menu "用户权限菜单树"
// @Router /menu/user-permission-tree [get]
func (m *MenuController) GetUserPermissionMenuTree(ctx *gin.Context) {
	// 从上下文中获取用户ID和用户名
	userId, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的用户ID"})
		return
	}

	tree, err := m.MenuService.GetUserPermissionMenuTree(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": tree})
}
