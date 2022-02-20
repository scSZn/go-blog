package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/service"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
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
		global.Logger.Errorf(ctx, "admin.CreateTag: create tag error, err: %v", err)
		if errors.Is(err, errcode.TagAlreadyExistError) {
			response.ReturnError(errcode.TagAlreadyExistError)
			return
		}
		response.ReturnError(errcode.CreateTagError)
		return
	}
	response.ReturnData("create tag success")
}
