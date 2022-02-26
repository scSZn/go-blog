package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/global"
)

func NewEnforcer() (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapterByDB(global.DB)
	if err != nil {
		return nil, err
	}
	enforcer, err := casbin.NewEnforcer(conf.GetAppSetting().CasbinModel, adapter)
	if err != nil {
		return nil, err
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return enforcer, nil
}
