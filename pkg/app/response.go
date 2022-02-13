package app

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/pkg/errcode"
	"net/http"
)

type Response struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		ctx: ctx,
	}
}

func (r *Response) ReturnJSON(data interface{}) {
	r.ctx.JSON(http.StatusOK, data)
}

func (r *Response) ReturnError(err *errcode.Error) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"code":    err.Code,
		"message": err.Message,
		"details": err.Detail,
	})
}
