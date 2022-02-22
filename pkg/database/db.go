package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/global"
)

func NewEngine(setting *conf.DatabaseSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		setting.Username,
		setting.Password,
		setting.Protocol,
		setting.Host,
		setting.Port,
		setting.Dbname,
		setting.Charset,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	db.Logger = logger.New(global.Logger, logger.Config{})
	if conf.GetEnv() == consts.EnvDev {
		db = db.Debug()
	}
	return db, nil
}
