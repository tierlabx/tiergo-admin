package crud

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ICrud[T any] interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	Page(*gin.Context)
}

type Crud[T any, CreateDTO any] struct {
	DB *gorm.DB
}

func (c Crud[T, CreateDTO]) Create(ctx *gin.Context) {
	var dto CreateDTO

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}
	var entity T

	err := copier.Copy(&entity, &dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "参数错误: " + err.Error()})
		return
	}
	if err := c.DB.Create(&entity).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "创建失败: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "创建成功", "data": entity})
}

func (c Crud[T, CreateDTO]) Update(ctx *gin.Context) {
	var entity T
	var dto CreateDTO

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的ID"})
		return
	}

	if err := c.DB.First(&entity, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}
	if err := copier.Copy(&entity, &dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "数据映射失败: " + err.Error()})
		return
	}
	if err := c.DB.Model(&entity).Updates(entity).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功", "data": entity})

}

func (c Crud[T, CreateDTO]) Delete(ctx *gin.Context) {
	var entity T
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的ID"})
		return
	}
	if err := c.DB.Delete(&entity, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "删除失败: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})

}

func (c Crud[T, CreateDTO]) Page(ctx *gin.Context) {
	var entity T
	var total int64
	var list []T
	// 分页参数
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if err := c.DB.Model(&entity).Count(&total).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	offset := (page - 1) * limit
	if err := c.DB.Model(&entity).Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取列表失败: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取列表成功",
		"data": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"data":  list,
		},
	})

}
