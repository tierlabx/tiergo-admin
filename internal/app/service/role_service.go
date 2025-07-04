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
