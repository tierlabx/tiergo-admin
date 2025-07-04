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

// 菜单权限树

// 递归构建树 // 指针允许null
func buildTreeMenu(menus []model.Menu, parentId *uint64) []model.Menu {
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
