package global

import (
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"

	"github.com/scSZn/blog/pkg/logger"
)

var (
	Logger   *logger.Logger
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
)
