package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/global"
)

func Default() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request
		global.Logger.Infof(ctx, "request coming, message: %+v", request)
		ctx.Next()
	}
}
