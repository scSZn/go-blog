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

func (r *Response) ReturnData(data interface{}) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.Code,
		"msg":  errcode.Success.Message,
		"data": data,
	})
}

func (r *Response) ReturnList(data interface{}, pager Pager, total int64) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.Code,
		"msg":  errcode.Success.Message,
		"data": gin.H{
			"total": total,
			"page":  pager.GetPage(),
			"limit": pager.GetLimit(),
			"list":  data,
		},
	})
}

func (r *Response) ReturnError(err error) {
	if newErr, ok := err.(*errcode.Error); ok {
		r.ctx.JSON(http.StatusOK, gin.H{
			"code":    newErr.Code,
			"msg":     newErr.Message,
			"details": newErr.Detail,
		})
		return
	}
	r.ctx.JSON(http.StatusInternalServerError, gin.H{})
}
