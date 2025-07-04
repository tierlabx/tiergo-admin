package casbin

import (
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
