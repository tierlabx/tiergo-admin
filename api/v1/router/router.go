package router

import (
	"tier-up/api/v1/controller"
	"tier-up/internal/app/middleware/auth"
	"tier-up/internal/app/middleware/jwt"
	"tier-up/internal/app/model"
	"tier-up/internal/app/service"
	"tier-up/internal/crud"

	_ "tier-up/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func SetupDigRouter(r *gin.Engine, c *dig.Container) error {
	return c.Invoke(func(
		db *gorm.DB,
		jwtService *jwt.JWTService,
		userService *service.UserService,
		roleService *service.RoleService,
		userController *controller.UserController,
		roleController *controller.RoleController,
		menuController *controller.MenuController,
	) {
		// 设置API路由组
		api := r.Group("/api/v1")
		api.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
		// 不需要认证的路由
		{
			api.POST("/register", userController.Register)
			api.POST("/login", userController.Login)
		}

		// 需要登录认证的路由
		authGroup := api.Group("")
		authGroup.Use(jwtService.JWTAuthMiddleware())
		{

			// 需要权限验证的路由
			rbacGroup := authGroup.Group("")
			rbacGroup.Use(auth.AuthMiddleware())
			{
				// 用户相关
				crud.RegisterCrudRoutes[model.User, model.UserReq](rbacGroup, db)
				rbacGroup.GET("/user/page", userController.Page)
				rbacGroup.GET("/user/info", userController.GetUserInfo)
				rbacGroup.PUT("/user/password", userController.ChangePassword)
				rbacGroup.POST("/user/update/:id", userController.Update)
				// 用户角色管理
				rbacGroup.POST("/user/:id/role", userController.AssignRole)

				// 角色管理
				crud.RegisterCrudRoutes[model.Role, model.RoleReq](rbacGroup, db)
				rbacGroup.GET("/role/:id", roleController.GetRoleByID)
				rbacGroup.POST("/role/permission-menu", roleController.AddPermissionMenu)
				rbacGroup.GET("/role/menu/:id", roleController.GetRoleMenu)

				// api权限管理
				rbacGroup.POST("/permission", roleController.AddPermission)
				rbacGroup.DELETE("/permission", roleController.RemovePermission)
				rbacGroup.GET("/role-permissions/:name", roleController.GetPermissions)

				// 菜单管理
				crud.RegisterCrudRoutes[model.Menu, model.MenuReq](rbacGroup, db)
				rbacGroup.GET("/menu/tree", menuController.GetMenuTree)
			}
		}
	})
}
