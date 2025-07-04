package controller

import (
	"net/http"
	"strconv"
	"tier-up/internal/app/service"

	"github.com/gin-gonic/gin"
)

// RoleController 角色控制器
type RoleController struct {
	RoleService *service.RoleService
}

// NewRoleController 创建角色控制器
func NewRoleController(roleService *service.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

// GetRoleByID 获取角色
// @Summary 获取角色详情
// @Description 根据ID获取角色详情
// @Tags Role
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "角色ID"
// @Success 200 {object} map[string]interface{} "角色详情"
// @Failure 400 {object} map[string]interface{} "无效的角色ID"
// @Failure 500 {object} map[string]interface{} "获取角色失败"
// @Router /role/{id} [get]
func (c *RoleController) GetRoleByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的角色ID"})
		return
	}

	role, err := c.RoleService.GetRoleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取角色失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取角色成功", "data": role})
}

// AddPermission 添加权限
// @Summary 添加权限
// @Description 为角色添加访问路径的权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body service.PermissionRequest true "权限信息"
// @Success 200 {object} map[string]interface{} "添加成功"
// @Failure 400 {object} map[string]interface{} "参数错误"
// @Failure 500 {object} map[string]interface{} "添加权限失败"
// @Router /permission [post]
func (c *RoleController) AddPermission(ctx *gin.Context) {
	var req service.PermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}

	if err := c.RoleService.AddPermission(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "添加权限失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "添加权限成功"})
}

// RemovePermission 移除权限
// @Summary 移除权限
// @Description 移除角色的访问路径权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body service.PermissionRequest true "权限信息"
// @Success 200 {object} map[string]interface{} "移除成功"
// @Failure 400 {object} map[string]interface{} "参数错误"
// @Failure 500 {object} map[string]interface{} "移除权限失败"
// @Router /permission [delete]
func (c *RoleController) RemovePermission(ctx *gin.Context) {
	var req service.PermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}

	if err := c.RoleService.RemovePermission(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "移除权限失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "移除权限成功"})
}

// GetPermissions 获取角色权限
// @Summary 获取角色权限
// @Description 获取指定角色的所有权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name path string true "角色名称"
// @Success 200 {object} map[string]interface{} "权限列表"
// @Failure 400 {object} map[string]interface{} "角色名称不能为空"
// @Failure 500 {object} map[string]interface{} "获取角色权限失败"
// @Router /role-permissions/{name} [get]
func (c *RoleController) GetPermissions(ctx *gin.Context) {
	roleName := ctx.Param("name")
	if roleName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "角色名称不能为空"})
		return
	}

	permissions, err := c.RoleService.GetPermissions(roleName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取角色权限失败: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取角色权限成功",
		"data":    permissions,
	})
}
