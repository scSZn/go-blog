package logger

import (
	"fmt"
	"github.com/scSZn/blog/consts"
	"github.com/sirupsen/logrus"
	"strings"
)

type LogFormatter struct {
}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	level := entry.Level
	time := entry.Time.Format(consts.LogTimeLayout)
	caller := entry.Caller.Function
	var slice = make([]string, 0, len(entry.Data)+1)
	for k, v := range entry.Data {
		slice = append(slice, fmt.Sprintf("%s=%v", k, v))
	}
	entry.Context.Value(consts.LogTraceKey)
	slice = append(slice, fmt.Sprintf("%s=%s", "message", entry.Message))
	return []byte(fmt.Sprintf("[%s][%s][%s] %s\n", level, time, caller, strings.Join(slice, "||"))), nil
}
