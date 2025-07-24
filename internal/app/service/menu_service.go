package service

import (
	"tier-up/internal/app/model"

	"gorm.io/gorm"
)

// RoleService 角色服务
type MenuService struct {
	DB *gorm.DB
}

// NewRoleService 创建角色服务
func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{
		DB: db,
	}
}

// 菜单树
func (m *MenuService) Tree() ([]model.Menu, error) {
	var menus []model.Menu
	if err := m.DB.Find(&menus).Error; err != nil {
		return nil, err
	}
	tree := buildTreeMenu(menus, nil)
	return tree, nil
}

// 递归构建树 // 指针允许null
func buildTreeMenu(menus []model.Menu, parentId *int) []model.Menu {
	var tree []model.Menu
	for _, m := range menus {
		if (m.ParentId == nil && parentId == nil) || (m.ParentId != nil && parentId != nil && *m.ParentId == *parentId) {
			children := buildTreeMenu(menus, &m.ID)
			m.Children = children
			tree = append(tree, m)
		}
	}
	return tree
}

// 获取用户所有的菜单权限
/* func (m *MenuService) GetUserPermissionMenuTree(userId uint) ([]model.Menu, error) {
	// 获取用户的所有角色
	var roles []model.Role
	if err := m.DB.Model(&model.User{}).Where("id = ?", userId).Preload("Role").Find(&roles).Error; err != nil {
		return nil, err
	}
	// 获取角色所有的菜单权限
	for _, role := range roles {
		permissions, err := m.GetPermissions(role.Name)
	}

} */
