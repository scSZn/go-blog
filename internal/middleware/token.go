package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
	"github.com/scSZn/blog/pkg/util"
)

const tokenHeaderKey = "token"

func TokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(tokenHeaderKey)
		if token == "" {
			token = ctx.Query(tokenHeaderKey)
		}

		response := app.NewResponse(ctx)
		if token == "" {
			global.Logger.Errorf(ctx, "token is empty")
			response.ReturnError(errcode.AuthorizationTokenInvalid)
			ctx.Abort()
			return
		}
		claims, err := util.ValidToken(token)
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch ve.Errors {
			case jwt.ValidationErrorExpired:
				global.Logger.Errorf(ctx, "token is expired, token is %+v", token)
				response.ReturnError(errcode.AuthorizationTokenTimeout)
				ctx.Abort()
				return
			default:
				global.Logger.Errorf(ctx, "token is invalid, token is %+v", token)
				response.ReturnError(errcode.AuthorizationTokenInvalid)
				ctx.Abort()
				return
			}
		}

		global.Logger.Infof(ctx, "user %+v coming", claims.Username)
		ctx.Set("uid", claims.Uid)
		ctx.Next()
	}
}
