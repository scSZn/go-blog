package global

import (
	"gorm.io/gorm"

	"github.com/scSZn/blog/pkg/logger"
)

var (
	Logger *logger.Logger
	DB     *gorm.DB
)
