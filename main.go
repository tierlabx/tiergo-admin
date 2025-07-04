package main

import (
	"fmt"

	"tier-up/api/v1/router"
	_ "tier-up/docs" // 导入swagger文档
	"tier-up/internal/app/middleware/casbin"
	"tier-up/internal/config"
	"tier-up/internal/db"
	"tier-up/internal/di"

	"github.com/gin-gonic/gin"
)

// @title           Tier Up API
// @version         1.0
// @description     Tier Up项目的API服务

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:88
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("|---------------------------|")
	fmt.Println("|----------admin------------|")
	fmt.Println("|---------------------------|")

	// 1. 初始化配置
	cfg := config.Load()

	// 2. 初始化数据库
	sqlDB, gormDB := db.InitDB(cfg)
	defer sqlDB.Close()

	// 3. 初始化Casbin
	casbin.InitCasbin(gormDB)

	// 4.依赖注入
	container := di.BuildContainer(gormDB)

	// 5.初始化 Gin 和路由
	r := gin.Default()
	router.SetupDigRouter(r, container)

	// 6. 启动服务器
	addr := fmt.Sprintf("%s:%s", cfg.WebApi.Host, cfg.WebApi.Port)
	fmt.Printf("服务器启动在 %s\n", addr)
	fmt.Println("Swagger文档地址: http://" + addr + "/api/v1/swagger/index.html")
	if err := r.Run(addr); err != nil {
		fmt.Printf("启动服务器失败: %v\n", err)
	}

	r.Run(":" + cfg.WebApi.Port)

}
