package model

// User 用户模型
type User struct {
	Base

	Username string `gorm:"size:50;not null;unique" json:"username"`
	Password string `gorm:"size:100;not null" json:"-"` // 密码不在JSON中返回
	Nickname string `gorm:"size:50" json:"nickname"`
	Email    string `gorm:"size:100;unique" json:"email"`
	Phone    string `gorm:"size:20" json:"phone"`
	Avatar   string `gorm:"size:255" json:"avatar"`
	Status   int    `gorm:"default:1" json:"status"` // 1:正常, 0:禁用

	Roles []Role `gorm:"many2many:user_roles;" json:"roles"`

	_ struct{} `crud:"prefix:/user,delete"`
}
type UserReq struct {
	Username string `json:"username"`
	Password string `gorm:"size:100;not null"` // 密码不在JSON中返回
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}

// UserRole 用户角色关联表
type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}
