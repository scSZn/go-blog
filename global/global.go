package global

import (
	"github.com/scSZn/blog/pkg/logger"
	"gorm.io/gorm"
)

var (
	Logger *logger.Logger
	DB     *gorm.DB
)
