package logger

import (
	"context"
	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/global"
	"github.com/sirupsen/logrus"
)

func NewLogger(ctx context.Context) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&LogFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(global.LogFileWriter)
	if conf.GetEnv() == consts.EnvDev {
		log.SetLevel(logrus.TraceLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	return log
}
