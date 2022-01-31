package model

import (
	"gorm.io/gorm"
	"time"
)

// Model gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey"`
	status    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsDel     bool
}
