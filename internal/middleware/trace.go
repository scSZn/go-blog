package middleware

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/consts"
	"time"

	"github.com/scSZn/blog/global"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var start = time.Now()

		request := ctx.Request
		ctx.Writer = &CustomResponseWriter{ResponseWriter: ctx.Writer, body: bytes.NewBufferString("")}
		requestString := fmt.Sprintf("%+v", request)
		ctx.Set(string(consts.LogTraceKey), genTraceId(requestString))

		global.Logger.Infof(ctx, map[string]interface{}{"request": requestString}, "request coming")

		ctx.Next()

		proc := float64(time.Since(start)) / 1e6
		writer, _ := ctx.Writer.(*CustomResponseWriter)
		global.Logger.Info(ctx, map[string]interface{}{
			"proc":     fmt.Sprintf("%fms", proc),
			"response": writer.GetResponseString(),
		})
	}
}

func genTraceId(key string) string {
	md5Instance := md5.New()
	_, err := md5Instance.Write([]byte(key))
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s%d", hex.EncodeToString(md5Instance.Sum(nil)), time.Now().UnixNano())
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (w *CustomResponseWriter) GetResponseString() string {
	return w.body.String()
}
