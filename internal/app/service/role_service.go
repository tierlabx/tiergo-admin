package service

import (
	"tier-up/internal/app/middleware/casbin"
	"tier-up/internal/app/model"

	"gorm.io/gorm"
)

// RoleService 角色服务
type RoleService struct {
	DB *gorm.DB
}

// PermissionRequest 权限请求
type PermissionRequest struct {
	Role   string `json:"role" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}
type PermissionMenuRequest struct {
	RoleId  int   `json:"role_id" binding:"required"`
	MenuIds []int `json:"menu_ids" binding:"required"`
}

// NewRoleService 创建角色服务
func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		DB: db,
	}
}

// GetRoleByID 通过ID获取角色
func (s *RoleService) GetRoleByID(id uint) (*model.Role, error) {
	var role model.Role
	if err := s.DB.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// AddPermission 添加权限
func (s *RoleService) AddPermission(req PermissionRequest) error {
	cs := casbin.GetInstance()
	_, err := cs.AddPolicy(req.Role, req.Path, req.Method)
	return err
}

// RemovePermission 移除权限
func (s *RoleService) RemovePermission(req PermissionRequest) error {
	cs := casbin.GetInstance()
	_, err := cs.RemovePolicy(req.Role, req.Path, req.Method)
	return err
}

// GetPermissions 获取角色的所有权限
func (s *RoleService) GetPermissions(roleName string) ([][]string, error) {
	cs := casbin.GetInstance()
	return cs.GetEnforcer().GetFilteredPolicy(0, roleName)
}

// 添加菜单权限
func (s *RoleService) AddPermissionMenu(roleId int, menuIds []int) error {
	var role model.Role
	if err := s.DB.First(&role, roleId).Error; err != nil {
		return err
	}
	var menus []model.Menu
	if err := s.DB.Where("id in ?", menuIds).Find(&menus).Error; err != nil {
		return err
	}
	// 添加角色关联
	if err := s.DB.Model(&role).Association("Menu").Append(&menus); err != nil {
		return err
	}
	return nil
}

// 获取角色的菜单ids
func (s *RoleService) GetRoleMenuIds(roleId int) ([]int, error) {
	var role model.Role
	if err := s.DB.First(&role, roleId).Error; err != nil {
		return nil, err
	}
	var menus []model.Menu
	if err := s.DB.Model(&role).Association("Menu").Find(&menus); err != nil {
		return nil, err
	}
	var menuIds []int
	for _, menu := range menus {
		menuIds = append(menuIds, menu.ID)
	}
	return menuIds, nil
}

// 获取角色的菜单
func (s *RoleService) GetRoleMenu(roleId int) ([]model.Menu, error) {
	var role model.Role
	if err := s.DB.First(&role, roleId).Error; err != nil {
		return nil, err
	}
	var menus []model.Menu
	if err := s.DB.Model(&role).Association("Menu").Find(&menus); err != nil {
		return nil, err
	}

	return menus, nil
}
