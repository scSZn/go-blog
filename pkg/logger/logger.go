package logger

import (
	"context"
	"fmt"
	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/consts"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime"
	"strings"
)

type Logger struct {
	*logrus.Logger
}

const depth = 2

func (l *Logger) Tracef(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Tracef(format, args...)
}

func (l *Logger) Debug(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Debug(args...)
}

func (l *Logger) Debugf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Debugf(format, args...)
}

func (l *Logger) Info(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Info(args...)
}

func (l *Logger) Infof(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Infof(format, args)
}

func (l *Logger) Warn(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Warn(args...)
}

func (l *Logger) Warnf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Warnf(format, args...)
}

func (l *Logger) Error(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Error(args...)
}

func (l *Logger) Errorf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Errorf(format, args...)
}

func (l *Logger) Fatal(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Fatal(args...)
}

func (l *Logger) Fatalf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	fields[consts.CallerKey] = getCallerInfo(depth)
	l.Logger.WithFields(fields).Fatalf(format, args...)
}

func addTraceInfo(ctx context.Context, fields map[string]interface{}) {
	traceId := ctx.Value(consts.LogTraceKey)
	fields[string(consts.LogTraceKey)] = traceId
}

func getCallerInfo(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s: %d", file, line)
}

// 对gorm的日志做处理，去除中间的换行符
func (l *Logger) Printf(format string, args ...interface{}) {
	format = strings.ReplaceAll(format, "\n", "")
	l.Logger.Printf(format, args...)
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
