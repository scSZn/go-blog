package logger

import (
	"context"
	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/consts"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

type Logger struct {
	*logrus.Logger
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//TODO implement me
	panic("implement me")
}

func (l *Logger) Tracef(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Tracef(format, args...)
}

func (l *Logger) Debug(ctx context.Context, args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

func (l *Logger) Info(ctx context.Context, args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *Logger) Warn(ctx context.Context, args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *Logger) Error(ctx context.Context, args ...interface{}) {
	l.Logger.Error(args...)
}

func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

func (l *Logger) Fatal(ctx context.Context, args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l *Logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

func NewLogger(setting *conf.LogSetting) (*Logger, error) {
	logger := &lumberjack.Logger{
		Filename:   setting.GetLogPath(),
		MaxSize:    setting.MaxSize,
		MaxBackups: setting.MaxBackups,
	}
	log := logrus.New()
	log.SetFormatter(&LogFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(logger)
	if conf.GetEnv() == consts.EnvDev {
		log.SetLevel(logrus.TraceLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	return &Logger{
		log,
	}, nil
}
