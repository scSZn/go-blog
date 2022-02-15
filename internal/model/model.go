package model

import (
	"time"

	"gorm.io/gorm"
)

// Model gorm.Model 的定义
type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Status    uint8          `json:"status" gorm:"default:1"` // 1：草稿，2：发布，3：禁用，4：删除
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	IsDel     bool           `json:"is_del" gorm:"default:false"`
}
