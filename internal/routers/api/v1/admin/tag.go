package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/pkg/errcode"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/service"
	"github.com/scSZn/blog/pkg/app"
)

func CreateTag(ctx *gin.Context) {
	request := service.CreateTagRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.CreateTag: bind error, err: %v", err)
		return
	}

	response := app.NewResponse(ctx)
	svc := service.NewTagService(ctx)
	err = svc.CreateTag(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	response.ReturnData("create tag success")
}

func DeleteTag(ctx *gin.Context) {
	request := service.DeleteTagRequest{}
	err := ctx.BindUri(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.CreateTag: bind error, err: %v", err)
		return
	}

	response := app.NewResponse(ctx)
	svc := service.NewTagService(ctx)
	err = svc.DeleteTag(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	response.ReturnData("delete tag success")
}

func ListTag(ctx *gin.Context) {
	request := service.ListTagRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.CreateTag: bind error, err: %v", err)
		return
	}

	response := app.NewResponse(ctx)
	svc := service.NewTagService(ctx)
	data, err := svc.ListTag(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	total, err := svc.CountTag(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	response.ReturnList(data, request.Pager, total)
}

func TagStatus(ctx *gin.Context) {
	tagStatus := conf.GetSetting().TagStatus
	response := app.NewResponse(ctx)

	response.ReturnData(tagStatus)
}

func TagStatusModify(ctx *gin.Context) {
	request := service.UpdateTagStatusRequest{}
	response := app.NewResponse(ctx)
	err := ctx.Bind(&request)
	if err != nil {
		response.ReturnError(errcode.BindError)
		global.Logger.Errorf(ctx, "admin.TagStatusModify: bind error, err: %v", err)
		return
	}
	err = ctx.BindUri(&request)
	if err != nil {
		response.ReturnError(errcode.BindError)
		global.Logger.Errorf(ctx, "admin.TagStatusModify: bind error, err: %v", err)
		return
	}

	svc := service.NewTagService(ctx)
	err = svc.UpdateTagStatus(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	response.ReturnData("update tag status success")
}
