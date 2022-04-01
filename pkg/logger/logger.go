package logger

import (
	"context"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/pkg/util"
)

type Fields = map[string]interface{}

type Logger interface {
	Tracef(context.Context, map[string]interface{}, string, ...interface{})
	Trace(context.Context, map[string]interface{}, ...interface{})
	Debugf(context.Context, map[string]interface{}, string, ...interface{})
	Debug(context.Context, map[string]interface{}, ...interface{})
	Infof(context.Context, map[string]interface{}, string, ...interface{})
	Info(context.Context, map[string]interface{}, ...interface{})
	Warnf(context.Context, map[string]interface{}, error, string, ...interface{})
	Warn(context.Context, map[string]interface{}, error, ...interface{})
	Errorf(context.Context, map[string]interface{}, error, string, ...interface{})
	Error(context.Context, map[string]interface{}, error, ...interface{})
	Fatalf(context.Context, map[string]interface{}, string, ...interface{})
	Fatal(context.Context, map[string]interface{}, ...interface{})
}

type DefaultLogger struct {
	*logrus.Logger
}

const depth = 2

func (l *DefaultLogger) Trace(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).Trace(args...)
}

func (l *DefaultLogger) Tracef(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).Tracef(format, args...)
}

func (l *DefaultLogger) Debug(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).Debug(args...)
}

func (l *DefaultLogger) Debugf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).Debugf(format, args...)
}

func (l *DefaultLogger) Info(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).Info(args...)
}

func (l *DefaultLogger) Infof(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).Infof(format, args)
}

func (l *DefaultLogger) Warn(ctx context.Context, fields map[string]interface{}, err error, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithError(err).WithFields(fields).Warn(args...)
}

func (l *DefaultLogger) Warnf(ctx context.Context, fields map[string]interface{}, err error, format string, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithError(err).WithFields(fields).Warnf(format, args...)
}

func (l *DefaultLogger) Error(ctx context.Context, fields map[string]interface{}, err error, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).WithError(err).Error(args...)
}

func (l *DefaultLogger) Errorf(ctx context.Context, fields map[string]interface{}, err error, format string, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).WithError(err).Errorf(format, args...)
}

func (l *DefaultLogger) Fatal(ctx context.Context, fields map[string]interface{}, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).Fatal(args...)
}

func (l *DefaultLogger) Fatalf(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	if _, ok := fields[consts.CallerKey]; !ok {
		fields[consts.CallerKey] = util.GetCallerFileAndLine(depth)
	}
	l.Logger.WithContext(ctx).WithFields(fields).Fatalf(format, args...)
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
