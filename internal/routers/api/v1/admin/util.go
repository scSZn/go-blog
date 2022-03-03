package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/pkg/cos"
	"github.com/scSZn/blog/pkg/errcode"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/pkg/app"
)

func UploadImage(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		global.Logger.Errorf(ctx, map[string]interface{}{
			"error": fmt.Sprintf("%+v", err),
		}, "bind error")
		response.ReturnError(errcode.UploadFail)
		return
	}

	filename := fileHeader.Filename
	if fileHeader.Size > consts.MaxUploadSize {
		global.Logger.Errorf(ctx, map[string]interface{}{
			"file": fileHeader,
		}, "upload file exceed max size")
		response.ReturnError(errcode.UploadFail)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		global.Logger.Errorf(ctx, map[string]interface{}{
			"file":  fileHeader,
			"error": err,
		}, "open file fail")
		response.ReturnError(errcode.UploadFail)
		return
	}

	location, err := cos.UploadImage(ctx, filename, file)
	if err != nil {
		global.Logger.Errorf(ctx, map[string]interface{}{
			"error": err,
			"file":  fileHeader,
		}, "upload file fail")
		response.ReturnError(errcode.UploadFail)
		return
	}

	response.ReturnData(location)
}
