package logger

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"

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

	// 获取TraceID
	traceId := getTraceInfo(entry.Context)
	var slice = make([]string, 0, len(entry.Data)+1)
	slice = append(slice, fmt.Sprintf("%s=%v", consts.LogTraceKey, traceId))

	// 拼装消息
	for k, v := range entry.Data {
		vStr, err := jsoniter.MarshalToString(v)
		if err != nil {
			return nil, err
		}
		slice = append(slice, fmt.Sprintf("%s=%s", k, vStr))
	}

	slice = append(slice, fmt.Sprintf("%s=%v", "message", entry.Message))
	return []byte(fmt.Sprintf("[%s][%s][%v] %s\n", level, time, caller, strings.Join(slice, "||"))), nil
}

// 如果是gin.Context，从gin.Context中获取traceId，否则从context中获取traceId
func getTraceInfo(ctx context.Context) interface{} {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(string(consts.LogTraceKey))
	}
	return ctx.Value(consts.LogTraceKey)
}
