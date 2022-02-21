package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Model gorm.Model 的定义
type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Status    uint8          `json:"status" gorm:"default:1"` // 1：草稿，2：发布，3：禁用，4：删除
	CreatedAt JsonTime       `json:"created_at" gorm:"column:created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt JsonTime       `json:"updated_at" gorm:"column:updated_at" time_format:"2006-01-02 15:04:05"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;index" time_format:"2006-01-02 15:04:05"`
	IsDel     bool           `json:"is_del" gorm:"default:false"`
}

// 定义个类型别名
type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 15:04:05")) // Format内即是你想转换的格式
	return []byte(stamp), nil
}
