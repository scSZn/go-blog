package database

import (
	"fmt"
	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewEngine(setting *conf.DatabaseSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s",
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
	return db, nil
}
