package conf

import (
	"github.com/spf13/viper"
	"log"
	"path"
	"sync"
	"time"
)

var setting *Setting
var once sync.Once

func GetSetting() *Setting {
	once.Do(func() {
		v := viper.New()
		v.SetConfigType("yml")
		v.AddConfigPath("./conf")
		v.SetConfigName("setting")

		err := v.ReadInConfig()
		if err != nil {
			log.Fatalf("read conf fail, err: %v", err)
		}
		err = v.Unmarshal(&setting)
		if err != nil {
			log.Fatalf("unmashal conf fail, err: %v", err)
		}
	})
	return setting
}

type Setting struct {
	Env string
	Log LogSetting
}

type LogSetting struct {
	Path     string
	Filename string
	Suffix   string
}

func GetEnv() string {
	return GetSetting().Env
}

func GetLogPath() string {
	setting := GetSetting()
	now := time.Now()
	suffix := now.Format(setting.Log.Suffix)
	return path.Join(setting.Log.Path, setting.Log.Filename) + "." + suffix
}
