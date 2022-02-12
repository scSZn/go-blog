package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/global"
	"io/ioutil"
)

func Default() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request
		args := make(map[string]interface{})
		args["method"] = request.Method
		args["path"] = request.URL.Path
		if args["path"] == "" {
			args["path"] = request.URL.RawPath
		}

		bodyData, err := ioutil.ReadAll(request.Body)
		if err != nil {
			global.Logger.Errorf(ctx, "get request body error, err: %v", err)
		} else {
			request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
		}

		marshal, err := json.Marshal(args)
		if err != nil {
			global.Logger.Errorf(ctx, "middle default json marshal error, err: %v", err)
			ctx.Next()
		}
		global.Logger.Infof(ctx, "request coming, message: %v", string(marshal))
	}
}
