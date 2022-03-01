package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/conf"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/service"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
)

// CreateArticle 新建博客
func CreateArticle(ctx *gin.Context) {
	request := service.CreateArticleRequest{}
	response := app.NewResponse(ctx)

	err := ctx.Bind(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.CreateArticle: bind error, err: %+v", err)
		response.ReturnError(errcode.ClientRequestError)
		return
	}

	svc := service.NewArticleService(ctx)
	err = svc.CreateArticle(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	response.ReturnData("create article success")
}

func ListArticleAdmin(ctx *gin.Context) {
	request := service.ListArticleRequest{}
	response := app.NewResponse(ctx)

	err := ctx.Bind(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "v1.ListArticleAdmin: bind error, err: %+v", err)
		response.ReturnError(errcode.ClientRequestError)
		return
	}

	svc := service.NewArticleService(ctx)
	total, err := svc.Count(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	articles, err := svc.List(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	response.ReturnList(articles, request.Pager, total)
}

func ArticleStatus(ctx *gin.Context) {
	articleStatus := conf.GetSetting().ArticleStatus
	response := app.NewResponse(ctx)
	response.ReturnData(articleStatus)
}
