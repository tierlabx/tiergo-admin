package model

// Role 角色模型
type Role struct {
	Base
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	DisplayName string `gorm:"size:100" json:"display_name"`
	Description string `gorm:"size:200" json:"description"`

	Menu []Menu `gorm:"many2many:role_menus;" json:"-"`

	_ struct{} `crud:"prefix:/role,create,update,delete,page"`
}

type RoleReq struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

// RoleMenu 角色菜单关联表
type RoleMenu struct {
	RoleId  uint `gorm:"primaryKey"`
	MenuIds uint
}
