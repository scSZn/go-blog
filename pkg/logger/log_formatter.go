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

	caller := entry.Data[consts.CallerKey]
	delete(entry.Data, consts.CallerKey)

	// 确保traceId放在前面
	traceKey := string(consts.LogTraceKey)
	var slice = make([]string, 0, len(entry.Data)+1)
	traceId := entry.Data[traceKey]
	slice = append(slice, fmt.Sprintf("%s=%v", consts.LogTraceKey, traceId))
	delete(entry.Data, traceKey)

	// 拼装消息
	for k, v := range entry.Data {
		slice = append(slice, fmt.Sprintf("%s=%v", k, v))
	}
	//entry.Context.Value(consts.LogTraceKey)
	slice = append(slice, fmt.Sprintf("%s=%s", "message", entry.Message))
	return []byte(fmt.Sprintf("[%s][%s][%v] %s\n", level, time, caller, strings.Join(slice, "||"))), nil
}
