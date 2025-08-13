package controller

import (
	"net/http"
	"strconv"
	"tier-up/internal/app/model"
	"tier-up/internal/app/service"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	UserService *service.UserService
}

// PasswordRequest 密码更新请求
type PasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=100"`
}

// RoleRequest 角色请求
type RoleRequest struct {
	RoleIDs []int `json:"role_ids" binding:"required"`
}

// NewUserController 创建用户控制器
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Register 用户注册
// @Summary 用户注册
// @Tags User
// @Accept json
// @Produce json
// @Param data body service.RegisterRequest true "用户注册信息"
// @Success 200 {object} Response[model.User] "注册成功"
// @Failure 400 {object} Response[any] "参数错误"
// @Failure 500 {object} Response[any] "注册失败"
// @Router /register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req service.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}

	user, err := c.UserService.Register(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "注册失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功", "data": user})
}

// Login 用户登录
// @Summary 用户登录
// @Tags User
// @Accept json
// @Produce json
// @Param data body service.LoginRequest true "用户登录信息"
// @Success 200 {object} Response[LoginResponse] "登录成功，返回token"
// @Failure 400 {object} Response[any] "参数错误"
// @Failure 401 {object} Response[any] "登录失败"
// @Router /login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var req service.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}

	accessToken, user, err := c.UserService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 401, "message": "登录失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data": gin.H{
			"accessToken": accessToken,
			// "refreshToken": refreshToken,
			"user": user,
		},
	})
}

// GetUserInfo 获取用户信息
// @Summary 获取当前用户信息
// @Description 获取已登录用户的详细信息
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} Response[model.User] "用户信息"
// @Failure 401 {object} Response[any] "未认证"
// @Failure 500 {object} Response[any] "获取用户信息失败"
// @Router /user/info [get]
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	userIDValue, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未认证"})
		return
	}

	userID := userIDValue.(int)
	user, err := c.UserService.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取用户信息失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取用户信息成功", "data": user})
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Description 修改当前用户的密码
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body PasswordRequest true "密码信息"
// @Success 200 {object} Response[any] "修改成功"
// @Router /user/password [put]
func (c *UserController) ChangePassword(ctx *gin.Context) {
	userIDValue, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未认证"})
		return
	}

	userID := userIDValue.(int)

	var req PasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}

	if err := c.UserService.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "修改密码失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改密码成功"})
}

// Page 获取用户分页
// @Summary 分页
// @Description
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page path  int true "页码"
// @Param limit path int true "当页条数"
// @Success 200 {object} Response[model.PageResult[model.User]] "用户分页"
// @Router /user/page [get]
func (c *UserController) Page(ctx *gin.Context) {
	var req model.PageLimitReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "分页参数错误"})
		return
	}
	println("page:", req.Page)

	list, err := c.UserService.Page(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取列表失败: " + err.Error()})
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    200,
			"message": "成功",
			"data":    list,
		},
	)
}

// Update 更新用户
// @Summary 更新用户
// @Description
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body model.UserReq true "修改参数"
// @Success 200 {object} Response[model.UserReq] "更新用户"
// @Router /user/update/:id [post]
func (c *UserController) Update(ctx *gin.Context) {
	var dto model.UserReq
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的ID"})
		return
	}
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}
	if err := c.UserService.UpdateFromDTO(int(id), &dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新失败: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功", "data": dto})
}

// AssignRole 分配角色
// @Summary 分配角色给用户
// @Description 为指定用户分配角色
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "用户ID"
// @Param data body RoleRequest true "角色信息"
// @Success 200 {object} Response[any] "分配成功"
// @Router /user/{id}/role [post]
func (c *UserController) AssignRole(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的用户ID"})
		return
	}

	var req RoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error(), "data": true})
		return
	}

	if err := c.UserService.AssignRoleToUser(int(userID), req.RoleIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "分配角色失败: " + err.Error(), "data": true})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "分配角色成功", "data": true})
}
