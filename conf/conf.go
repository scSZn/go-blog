package conf

import (
	"log"
	"path"
	"sync"
	"time"

	"github.com/spf13/viper"

	"github.com/scSZn/blog/consts"
)

var setting *Setting
var once sync.Once

func init() {
	once.Do(func() {
		v := viper.New()
		v.SetConfigType("yml")
		v.AddConfigPath("./conf")
		v.SetConfigName("setting")
		v.RegisterAlias("tag_status", "status.tag")
		v.RegisterAlias("article_status", "status.article")

		err := v.ReadInConfig()
		if err != nil {
			log.Fatalf("read conf fail, err: %v", err)
		}

		v.SetConfigName("secret")
		err = v.MergeInConfig()
		if err != nil {
			log.Fatalf("merge conf fail, err: %v", err)
		}
		err = v.Unmarshal(&setting)
		if err != nil {
			log.Fatalf("unmashal conf fail, err: %v", err)
		}
	})
}

func GetSetting() *Setting {
	return setting
}

type Setting struct {
	Env           string           `json:"env" mapstructure:"env"`
	AppSetting    *AppSetting      `json:"app_setting" mapstructure:"app"`
	LogSetting    *LogSetting      `json:"log_setting" mapstructure:"log"`
	DBSetting     *DatabaseSetting `json:"db_setting" mapstructure:"database"`
	COSSetting    *COSSetting      `json:"cos" mapstructure:"cos"`
	TagStatus     []Status         `json:"tag_status" mapstructure:"tag_status"`
	ArticleStatus []Status         `json:"article_status" mapstructure:"article_status"`
}

type AppSetting struct {
	Host         string   `json:"host" mapstructure:"host"`
	Port         string   `json:"port" mapstructure:"port"`
	CasbinModel  string   `json:"casbin_model" mapstructure:"casbin_model"`
	AllowOrigins []string `json:"allow_origins" mapstructure:"allow_origins"`
}

type LogSetting struct {
	Path       string `json:"path" mapstructure:"path"`
	Filename   string `json:"filename" mapstructure:"filename"`
	Suffix     string `json:"suffix" mapstructure:"suffix"`
	MaxSize    int    `json:"max_size" mapstructure:"max_size"`
	MaxBackups int    `json:"max_backups" mapstructure:"max_backups"`
}

type DatabaseSetting struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Dbname   string `json:"dbname" mapstructure:"dbname"`
	Charset  string `json:"charset" mapstructure:"charset"`
	Protocol string `json:"protocol" mapstructure:"protocol"`
}

type COSSetting struct {
	AppID     string `json:"app_id" mapstructure:"app_id"`
	SecretID  string `json:"secret_id" mapstructure:"secret_id"`
	SecretKey string `json:"secret_key" mapstructure:"secret_key"`
	BaseURL   string `json:"base_url" mapstructure:"base_url"`
}

type Status struct {
	Value   int    `json:"value" mapstructure:"value"`
	Name    string `json:"name" mapstructure:"name"`
	Display string `json:"display" mapstructure:"display"`
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

func GetAppSetting() *AppSetting {
	return setting.AppSetting
}

func GetListenAddr() string {
	appSetting := GetAppSetting()
	return appSetting.Host + ":" + appSetting.Port
}
