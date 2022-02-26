package admin

import (
	"github.com/gin-gonic/gin"

	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/service"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
)

func CreateTag(ctx *gin.Context) {
	request := service.CreateTagRequest{}
	response := app.NewResponse(ctx)

	err := ctx.Bind(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.CreateTag: bind error, err: %+v", err)
		response.ReturnError(errcode.ClientRequestError)
		return
	}

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
	response := app.NewResponse(ctx)

	err := ctx.BindUri(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.CreateTag: bind error, err: %+v", err)
		response.ReturnError(errcode.ClientRequestError)
		return
	}

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
	response := app.NewResponse(ctx)

	err := ctx.BindQuery(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.CreateTag: bind error, err: %+v", err)
		response.ReturnError(errcode.ClientRequestError)
		return
	}

	svc := service.NewTagService(ctx)
	data, total, err := svc.ListTag(&request)
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

func UpdateTag(ctx *gin.Context) {
	request := service.UpdateTagRequest{}
	response := app.NewResponse(ctx)
	// 绑定JSON
	err := ctx.Bind(&request)
	if err != nil {
		response.ReturnError(errcode.ClientRequestError)
		global.Logger.Errorf(ctx, "admin.UpdateTag: bind error, err: %+v", err)
		return
	}
	// 绑定路径参数
	err = ctx.BindUri(&request)
	if err != nil {
		response.ReturnError(errcode.ClientRequestError)
		global.Logger.Errorf(ctx, "admin.UpdateTag: bind error, err: %+v", err)
		return
	}

	svc := service.NewTagService(ctx)
	err = svc.UpdateTag(&request)
	if err != nil {
		response.ReturnError(err)
		return
	}
	response.ReturnData("update tag status success")
}
