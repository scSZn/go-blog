package model

import (
	"time"

	"gorm.io/gorm"
)

// Model gorm.Model 的定义
type Model struct {
	ID        uint  `gorm:"primaryKey"`
	Status    uint8 `gorm:"default:1"` // 1：草稿，2：发布，3：禁用，4：删除
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsDel     bool           `gorm:"default:false"`
}
