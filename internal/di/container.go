package di

import (
	"tier-up/api/v1/controller"
	"tier-up/internal/app/middleware/jwt"
	"tier-up/internal/app/service"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func BuildContainer(db *gorm.DB) *dig.Container {
	container := dig.New()

	// 基础服务
	container.Provide(func() *gorm.DB { return db })
	container.Provide(jwt.NewJWTService)

	// 业务服务
	container.Provide(service.NewUserService)
	container.Provide(service.NewRoleService)
	container.Provide(service.NewMenuService)
	// 控制器
	container.Provide(controller.NewUserController)
	container.Provide(controller.NewRoleController)
	container.Provide(controller.NewMenuController)

	return container
}
