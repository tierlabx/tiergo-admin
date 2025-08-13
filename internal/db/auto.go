package db

import (
	"fmt"
	"tier-up/internal/app/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		// 系统表
		&model.User{},
		&model.Role{},
		&model.UserRole{},
		&model.Menu{},
	); err != nil {
		panic(err)
	}
	var count int64
	db.Model(&model.Role{}).Where("name = ?", "super_admin").Count(&count)
	if count == 0 {
		role := model.Role{
			Name:        "super_admin",
			DisplayName: "超级管理员",
			Description: "拥有系统所有权限",
		}
		if err := db.Create(&role).Error; err != nil {
			fmt.Println("创建超级管理员失败，请手动创建")
		}
		// 密码加密
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("密码加密失败", err)
		}
		admin := model.User{
			Username: "admin",
			Password: string(hashedPassword),
			Nickname: "超级管理员",
			Email:    "admin@example.com",
			Phone:    "12345678901",
			Avatar:   "https://fastly.picsum.photos/id/646/640/480.jpg?hmac=4ilW8ljWVx1voBCCves3xCOhrsDC0ag5tBcz4wlK_Ls",
			Status:   1,
		}
		if err := db.Create(&admin).Error; err != nil {
			fmt.Println("创建超级管理员失败，请手动创建")
		}
		db.Model(&admin).Association("Roles").Append(&role)

	}

}
