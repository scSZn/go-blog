package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/scSZn/blog/pkg/logger"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
	"github.com/scSZn/blog/pkg/util"
)

const tokenHeaderKey = "token"

func PermissionVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(tokenHeaderKey)
		if token == "" {
			token = ctx.Query(tokenHeaderKey)
		}

		response := app.NewResponse(ctx)
		if token == "" {
			global.Logger.Errorf(ctx, nil, nil, "token is empty")
			response.ReturnError(errcode.AuthorizationTokenInvalid)
			ctx.Abort()
			return
		}
		claims, err := util.ValidToken(token)
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch ve.Errors {
			case jwt.ValidationErrorExpired:
				global.Logger.Errorf(ctx, logger.Fields{"token": token}, nil, "token is expired")
				response.ReturnError(errcode.AuthorizationTokenTimeout)
				ctx.Abort()
				return
			default:
				global.Logger.Errorf(ctx, logger.Fields{"token": token}, nil, "token is invalid")
				response.ReturnError(errcode.AuthorizationTokenInvalid)
				ctx.Abort()
				return
			}
		}

		global.Logger.Infof(ctx, nil, "user [%+v] attempt to access [%+v]", claims.Username, ctx.Request.URL.Path)

		enforce, err := global.Enforcer.Enforce(claims.Username, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil || !enforce {
			global.Logger.Errorf(ctx, nil, err, "user [%+v] not permission access [%+v]", claims.Username, ctx.Request.URL.Path)
			response.ReturnError(errcode.AuthorizationNotPermission)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
