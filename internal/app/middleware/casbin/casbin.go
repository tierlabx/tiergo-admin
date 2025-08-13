package casbin

import (
	"fmt"
	"log"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// CasbinService 权限管理服务
type CasbinService struct {
	Enforcer *casbin.Enforcer
	once     sync.Once
}

var (
	casbinService = &CasbinService{}
)

// InitCasbin 初始化Casbin
func InitCasbin(db *gorm.DB) *CasbinService {
	casbinService.once.Do(func() {
		// 使用gorm适配器
		adapter, err := gormadapter.NewAdapterByDB(db)
		if err != nil {
			log.Fatalf("初始化Casbin适配器失败: %v", err)
		}

		// 从文件加载模型配置
		m, err := model.NewModelFromFile("internal/app/middleware/casbin/model.conf")
		if err != nil {
			log.Fatalf("加载Casbin模型配置失败: %v", err)
		}

		// 创建Enforcer
		enforcer, err := casbin.NewEnforcer(m, adapter)
		if err != nil {
			log.Fatalf("创建Casbin Enforcer失败: %v", err)
		}

		// 加载策略
		if err := enforcer.LoadPolicy(); err != nil {
			log.Fatalf("加载Casbin策略失败: %v", err)
		}

		casbinService.Enforcer = enforcer
	})

	return casbinService
}

// 初始化admin 角色权限
func (cs *CasbinService) InitAdmin() error {
	// 检查是否已存在策略 p, super_admin, *, *
	exists, err := cs.Enforcer.HasPolicy("super_admin", "*", "*")
	if err != nil {
		return fmt.Errorf("检查策略时出错: %v", err)
	}

	// 如果该策略已经存在，返回不需要添加
	if exists {
		fmt.Println("超级管理员权限已经存在")
		return nil
	}

	userId := "1"
	_, err_role := cs.AddRoleForUser(userId, "super_admin")
	if err_role != nil {
		return err_role
	}
	// p, super_admin, *, * 代表超级管理员可以访问所有资源并执行所有操作
	_, cs_err := cs.AddPolicy("super_admin", "*", "*")
	if cs_err != nil {
		return err
	}
	return nil
}

// GetEnforcer 获取Enforcer实例
func (cs *CasbinService) GetEnforcer() *casbin.Enforcer {
	return cs.Enforcer
}

// AddPolicy 添加策略
func (cs *CasbinService) AddPolicy(sub, obj, act string) (bool, error) {
	return cs.Enforcer.AddPolicy(sub, obj, act)
}

// RemovePolicy 删除策略
func (cs *CasbinService) RemovePolicy(sub, obj, act string) (bool, error) {
	return cs.Enforcer.RemovePolicy(sub, obj, act)
}

// AddRoleForUser 为用户添加角色
func (cs *CasbinService) AddRoleForUser(user, role string) (bool, error) {
	return cs.Enforcer.AddGroupingPolicy(user, role)
}

// DeleteRoleForUser 删除用户的角色
func (cs *CasbinService) DeleteRoleForUser(user, role string) (bool, error) {
	return cs.Enforcer.RemoveGroupingPolicy(user, role)
}

// Enforce 检查权限
func (cs *CasbinService) Enforce(sub, obj, act string) (bool, error) {
	return cs.Enforcer.Enforce(sub, obj, act)
}

// GetInstance 获取CasbinService实例
func GetInstance() *CasbinService {
	return casbinService
}
