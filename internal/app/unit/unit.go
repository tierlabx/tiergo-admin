package unit

import "tier-up/internal/app/model"

// 递归构建树 // 指针允许null
func BuildTreeMenu(menus []model.Menu, parentId *int) []model.Menu {
	var tree []model.Menu
	for _, m := range menus {
		if (m.ParentId == nil && parentId == nil) || (m.ParentId != nil && parentId != nil && *m.ParentId == *parentId) {
			children := BuildTreeMenu(menus, &m.ID)
			m.Children = children
			tree = append(tree, m)
		}
	}
	return tree
}

type HasID interface {
	GetID() int
}

// 切片去重
func UniqueStructByID[T HasID](entities []T) []T {
	m := make(map[int]struct{})
	var result []T
	for _, u := range entities {
		if _, exists := m[u.GetID()]; !exists {
			m[u.GetID()] = struct{}{}
			result = append(result, u)
		}
	}
	return result
}
