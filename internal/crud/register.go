package crud

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 根据配置 创建
func RegisterCrudRoutes[T any, C any](
	r *gin.RouterGroup,
	db *gorm.DB,
) {
	// 解析model tag配置
	config := ParseModelConfig[T]()
	var handle ICrud[T] = Crud[T, C]{DB: db}
	group := r.Group(config.Prefix)
	// 按需注册路由
	if config.Create {
		group.POST("/create", handle.Create)
	}
	if config.Update {
		group.PUT("/update/:id", handle.Update)
	}
	if config.Delete {
		group.DELETE("/delete/:id", handle.Delete)
	}
	if config.Page {
		group.GET("/page", handle.Page)
	}
}
