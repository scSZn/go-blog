package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/service"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
)

func CreateArticle(ctx *gin.Context) {
	request := service.CreateArticleRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "bind error, err: %v", err)
		return
	}

	response := app.NewResponse(ctx)
	svc := service.NewArticleService(ctx)
	err = svc.CreateArticle(request)
	if err != nil {
		global.Logger.Errorf(ctx, "create article error, err: %v", err)
		response.ReturnError(errcode.CreateArticleError)
		return
	}
	ctx.JSON(200, "create article success")
}
