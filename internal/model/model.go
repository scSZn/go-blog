package model

import (
	"time"

	"gorm.io/gorm"
)

// Model gorm.Model 的定义
type Model struct {
	ID        uint  `gorm:"primaryKey"`
	status    uint8 `gorm:"default:0"` // 0表示可用，1表示不可用
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsDel     bool           `gorm:"default:false"`
}
