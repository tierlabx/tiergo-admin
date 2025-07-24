package model

type Menu struct {
	Base
	Code      string `json:"code"`
	Name      string `json:"name" gorm:"not null;"`
	Path      string `json:"path" gorm:"not null;comment:api路径;"`
	Component string `json:"component" gorm:"not null;comment:组件路径;" `
	Icon      string `json:"icon" gorm:"comment:icon图标;"`
	Note      string `json:"note" gorm:"comment:备注;"`
	Type      int    `json:"type"`
	Status    *int   `json:"status" gorm:"comment:状态:1正常 2禁用;" `
	Sort      int    `json:"sort" gorm:"comment:显示顺序;"`
	ParentId  *int   `json:"parent_id" gorm:"column:parent_id"` // 允许为空的父ID

	Children []Menu `json:"children" gorm:"foreignKey:ParentId"`

	_ struct{} `crud:"prefix:/menu,create,update,delete"`
}

type MenuReq struct {
	Code      string `json:"code"`
	Name      string `json:"name" `
	Path      string `json:"path"  binding:"required" `
	Component string `json:"component"`
	Icon      string `json:"icon" `
	Note      string `json:"note" `
	Type      int    `json:"type"  binding:"required"`
	Status    *int   `json:"status"  `
	Sort      *int   `json:"sort"`
	ParentId  *int   `json:"parent_id" `
}

func (m Menu) GetID() int {
	return m.ID
}
