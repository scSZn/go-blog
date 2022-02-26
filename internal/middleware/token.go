package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

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
			response.ReturnError(errcode.AuthorizationTokenInvalid)
			ctx.Abort()
			return
		}
		claims, err := util.ValidToken(token)
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch ve.Errors {
			case jwt.ValidationErrorExpired:
				response.ReturnError(errcode.AuthorizationTokenTimeout)
				ctx.Abort()
				return
			default:
				response.ReturnError(errcode.AuthorizationTokenInvalid)
				ctx.Abort()
				return
			}
		}

		ctx.Set("uid", claims.Uid)
		ctx.Next()
	}
}
