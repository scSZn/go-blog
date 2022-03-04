package logger

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/pkg/util"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Tracef(context.Context, map[string]interface{}, string, ...interface{})
	Trace(context.Context, map[string]interface{}, ...interface{})
	Debugf(context.Context, map[string]interface{}, string, ...interface{})
	Debug(context.Context, map[string]interface{}, ...interface{})
	Infof(context.Context, map[string]interface{}, string, ...interface{})
	Info(context.Context, map[string]interface{}, ...interface{})
	Warnf(context.Context, map[string]interface{}, string, ...interface{})
	Warn(context.Context, map[string]interface{}, ...interface{})
	Errorf(context.Context, map[string]interface{}, string, ...interface{})
	Error(context.Context, map[string]interface{}, ...interface{})
	Fatalf(context.Context, map[string]interface{}, string, ...interface{})
	Fatal(context.Context, map[string]interface{}, ...interface{})
}

type DefaultLogger struct {
	*logrus.Logger
}

const depth = 2

func (l *DefaultLogger) Trace(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Trace(args...)
}

func (l *DefaultLogger) Tracef(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Tracef(format, args...)
}

func (l *DefaultLogger) Debug(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Debug(args...)
}

func (l *DefaultLogger) Debugf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Debugf(format, args...)
}

func (l *DefaultLogger) Info(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Info(args...)
}

func (l *DefaultLogger) Infof(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Infof(format, args)
}

func (l *DefaultLogger) Warn(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Warn(args...)
}

func (l *DefaultLogger) Warnf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Warnf(format, args...)
}

func (l *DefaultLogger) Error(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Error(args...)
}

func (l *DefaultLogger) Errorf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Errorf(format, args...)
}

func (l *DefaultLogger) Fatal(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Fatal(args...)
}

func (l *DefaultLogger) Fatalf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	addTraceInfo(ctx, fields)
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithFields(fields).Fatalf(format, args...)
}

func addTraceInfo(ctx context.Context, fields map[string]interface{}) {
	traceId := ctx.Value(consts.LogTraceKey)
	if ginCtx, ok := ctx.(*gin.Context); traceId == nil && ok {
		traceId = ginCtx.GetString(string(consts.LogTraceKey))
	}
	fields[string(consts.LogTraceKey)] = traceId
}

func NewLogger(setting *conf.LogSetting) (Logger, error) {
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

	return &DefaultLogger{
		log,
	}, nil
}
