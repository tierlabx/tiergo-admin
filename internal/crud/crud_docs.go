package crud


import (
  "github.com/gin-gonic/gin"
"tier-up/internal/app/model"
)

// ===== Auto-generated stub for Menu =====

// @Summary 创建 Menu
// @Description 创建 Menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param data body model.MenuReq true "Menu 数据"
// @Success 200 {object} MenuResponse
// @Router /menu/create [post]
func MenuCreateDoc(ctx *gin.Context) {}


// @Summary 删除 Menu
// @Description 删除 Menu
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {object} MenuResponse
// @Router /menu/delete/:id [delete]
func MenuDeleteDoc(ctx *gin.Context) {}


// @Summary 更新 Menu
// @Description 更新 Menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param data body model.MenuReq true "Menu 数据"
// @Success 200 {object} MenuResponse
// @Router /menu/update/:id [put]
func MenuUpdateDoc(ctx *gin.Context) {}


type MenuResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data   model. Menu `json:"data"`
}

type MenuPageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data struct {
		Page  int     `json:"page"`
		Limit int     `json:"limit"`
		Total int64   `json:"total"`
		Data  []model.Menu `json:"data"`
	} `json:"data"`
}

// ===== Auto-generated stub for Role =====

// @Summary 创建 Role
// @Description 创建 Role
// @Tags Role
// @Accept json
// @Produce json
// @Param data body model.RoleReq true "Role 数据"
// @Success 200 {object} RoleResponse
// @Router /role/create [post]
func RoleCreateDoc(ctx *gin.Context) {}


// @Summary 删除 Role
// @Description 删除 Role
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {object} RoleResponse
// @Router /role/delete/:id [delete]
func RoleDeleteDoc(ctx *gin.Context) {}


// @Summary 更新 Role
// @Description 更新 Role
// @Tags Role
// @Accept json
// @Produce json
// @Param data body model.RoleReq true "Role 数据"
// @Success 200 {object} RoleResponse
// @Router /role/update/:id [put]
func RoleUpdateDoc(ctx *gin.Context) {}


// @Summary 分页查询 Role
// @Description 分页查询 Role
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {object} RolePageResponse
// @Router /role/page [get]
func RolePageDoc(ctx *gin.Context) {}


type RoleResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data   model. Role `json:"data"`
}

type RolePageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data struct {
		Page  int     `json:"page"`
		Limit int     `json:"limit"`
		Total int64   `json:"total"`
		Data  []model.Role `json:"data"`
	} `json:"data"`
}

// ===== Auto-generated stub for User =====

// @Summary 删除 User
// @Description 删除 User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse
// @Router /user/delete/:id [delete]
func UserDeleteDoc(ctx *gin.Context) {}


type UserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data   model. User `json:"data"`
}

type UserPageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data struct {
		Page  int     `json:"page"`
		Limit int     `json:"limit"`
		Total int64   `json:"total"`
		Data  []model.User `json:"data"`
	} `json:"data"`
}
