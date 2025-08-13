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

		// 获取请求的路径和方法
		obj := c.Request.URL.Path
		act := c.Request.Method

		// 获取Casbin服务实例
		cs := casbin.GetInstance()

		// 使用用户角色作为主体进行权限检查
		sub := strconv.Itoa(userID.(int))
		roles, roleErr := cs.GetEnforcer().GetRolesForUser(sub)

		if roleErr != nil || len(roles) == 0 {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "用户未分配角色"})
			c.Abort()
			return
		}
		fmt.Printf("用户 %s 的角色列表: %v\n", sub, roles)
		// 使用用户的角色进行权限检查
		hasPermission := false
		for _, role := range roles {
			ok, err := cs.Enforce(role, obj, act)
			fmt.Printf("检查 %v， 角色 %v , 路径%v , 操作 %v", ok, role, obj, act)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "权限检查失败"})
				c.Abort()
				return
			}
			if ok {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "没有足够的权限"})
			c.Abort()
			return
		}

		c.Next()
	}
}
