package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/pkg/app"
	"strings"
)

const originHeaderKey = "Origin"
const accessControlAllowOrigin = "Access-Control-Allow-Origin"
const accessControlAllowMethods = "Access-Control-Allow-Methods"
const accessControlAllowCredentials = "Access-Control-Allow-Credentials"
const accessControlExposeHeaders = "Access-Control-Expose-Headers"
const accessControlAllowHeaders = "Access-Control-Allow-Headers"

func CORS() gin.HandlerFunc {
	allowOrigins := conf.GetAppSetting().AllowOrigins
	return func(ctx *gin.Context) {
		origin := ctx.GetHeader(originHeaderKey)
		for _, allowOrigin := range allowOrigins {
			if strings.Contains(allowOrigin, origin) {
				ctx.Header(accessControlAllowOrigin, origin)
				ctx.Header(accessControlAllowMethods, "GET, PUT, OPTIONS, POST, DELETE")
				ctx.Header(accessControlAllowCredentials, "true")
				ctx.Header(accessControlAllowHeaders, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, token")
				ctx.Header(accessControlExposeHeaders, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, token")
			}
		}
		if ctx.Request.Method == "OPTIONS" {
			response := app.NewResponse(ctx)
			response.ReturnData("Success")
			return
		}

		ctx.Next()
	}
}
