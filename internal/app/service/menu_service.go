package service

import (
	"tier-up/internal/app/model"
	"tier-up/internal/app/unit"

	"gorm.io/gorm"
)

// RoleService 角色服务
type MenuService struct {
	DB          *gorm.DB
	RoleService *RoleService
}

// NewRoleService 创建角色服务
func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{
		DB:          db,
		RoleService: NewRoleService(db),
	}
}

// 菜单树
func (m *MenuService) Tree() ([]model.Menu, error) {
	var menus []model.Menu
	if err := m.DB.Find(&menus).Error; err != nil {
		return nil, err
	}
	tree := unit.BuildTreeMenu(menus, nil)
	return tree, nil
}

// 获取用户所有的菜单权限
func (m *MenuService) GetUserPermissionMenuTree(userId int) ([]model.Menu, error) {
	// 获取用户的所有角色
	var user model.User
	if err := m.DB.Preload("Roles").First(&user, userId).Error; err != nil {
		return nil, err
	}
	// 获取所有角色所有的菜单权限
	var menus []model.Menu
	for _, role := range user.Roles {
		permissions, err := m.RoleService.GetRoleMenu(role.ID)
		if err != nil {
			return nil, err
		}
		menus = append(menus, permissions...)
	}
	// 去重
	menus = unit.UniqueStructByID(menus)
	tree := unit.BuildTreeMenu(menus, nil)
	return tree, nil
}
