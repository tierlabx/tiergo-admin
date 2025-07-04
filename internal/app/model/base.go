package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type PageLimitReq struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type PageResult[T any] struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
	List  []T   `json:"list"`
}
