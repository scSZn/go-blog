package main

import (
	"github.com/scSZn/blog/pkg/casbin"
	"log"

	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/pkg/database"
	"github.com/scSZn/blog/pkg/logger"
)

func Init() {
	var err error
	global.Logger, err = logger.NewLogger(conf.GetLogSetting())
	if err != nil {
		log.Fatalf("[main.Init] logger init fail, err: %v", err)
	}

	global.DB, err = database.NewEngine(conf.GetDatabaseSetting())
	if err != nil {
		log.Fatalf("[main.Init] db init fail, err: %v", err)
	}

	global.Enforcer, err = casbin.NewEnforcer()
	if err != nil {
		log.Fatalf("[main.Init] casbin enforcer init fail, err: %v", err)
	}

}
