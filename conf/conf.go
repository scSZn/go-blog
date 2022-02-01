package conf

import (
	"github.com/scSZn/blog/consts"
	"github.com/spf13/viper"
	"log"
	"path"
	"sync"
	"time"
)

var setting = GetSetting()
var once sync.Once

func GetSetting() *Setting {
	once.Do(func() {
		v := viper.New()
		v.SetConfigType("yml")
		v.AddConfigPath("./conf")
		v.SetConfigName("setting")
		v.RegisterAlias("dbsetting", "database")
		v.RegisterAlias("logsetting", "log")

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
	Env        string
	LogSetting *LogSetting
	DBSetting  *DatabaseSetting
}

type LogSetting struct {
	Path       string
	Filename   string
	Suffix     string
	MaxSize    int
	MaxBackups int
}

type DatabaseSetting struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Charset  string
	Protocol string
}

func GetEnv() string {
	if setting.Env == "" {
		return consts.EnvDev
	}
	return setting.Env
}

func (ls *LogSetting) GetLogPath() string {
	now := time.Now()
	suffix := now.Format(ls.Suffix)
	return path.Join(ls.Path, ls.Filename) + "." + suffix
}

func GetDatabaseSetting() *DatabaseSetting {
	return setting.DBSetting
}

func GetLogSetting() *LogSetting {
	return setting.LogSetting
}
