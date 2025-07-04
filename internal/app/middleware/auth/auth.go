package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"tier-up/internal/app/middleware/casbin"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 权限验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户ID和用户名
		userID, exists := c.Get("userID")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权的访问"})
			c.Abort()
			return
		}
		username, exists1 := c.Get("username")
		fmt.Println("用户 =", username, exists1)
		if username == "admin" {
			c.Next()
			return
		}
		// 获取请求的路径和方法
		obj := c.Request.URL.Path
		act := c.Request.Method

		// 获取Casbin服务实例
		cs := casbin.GetInstance()

		// 使用用户ID作为主体进行权限检查
		sub := strconv.Itoa(int(userID.(uint64)))
		ok, err := cs.Enforce(sub, obj, act)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "权限检查失败"})
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "没有足够的权限"})
			c.Abort()
			return
		}

		c.Next()
	}
}
