package logger

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/scSZn/blog/consts"
)

type LogFormatter struct {
}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	level := strings.ToUpper(entry.Level.String())
	time := entry.Time.Format(consts.LogTimeLayout)
	var slice = make([]string, 0, len(entry.Data)+1)
	for k, v := range entry.Data {
		slice = append(slice, fmt.Sprintf("%s=%v", k, v))
	}
	//entry.Context.Value(consts.LogTraceKey)
	slice = append(slice, fmt.Sprintf("%s=%s", "message", entry.Message))
	return []byte(fmt.Sprintf("[%s][%s] %s\n", level, time, strings.Join(slice, "||"))), nil
}
