package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/pkg/cos"
	"github.com/scSZn/blog/pkg/errcode"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/pkg/app"
)

func UploadImage(c *gin.Context) {
	response := app.NewResponse(c)
	fileHeader, err := c.FormFile("file")
	if err != nil {
		global.Logger.Errorf(c, "v1Admin.UploadImage: get file fail, err: %+v", err)
		response.ReturnError(errcode.UploadFail)
		return
	}

	filename := fileHeader.Filename
	if fileHeader.Size > consts.MaxUploadSize {
		global.Logger.Errorf(c, "v1Admin.UploadImage: upload file exceed max size, filename: %+v", filename)
		response.ReturnError(errcode.UploadFail)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		global.Logger.Errorf(c, "v1Admin.UploadImage: open file fail, err: %+v", err)
		response.ReturnError(errcode.UploadFail)
		return
	}

	location, err := cos.UploadImage(c, filename, file)
	if err != nil {
		global.Logger.Errorf(c, "v1Admin.UploadImage: upload file fail, filename is %+v, err: %+v", filename, err)
		response.ReturnError(errcode.UploadFail)
		return
	}

	response.ReturnData(location)
}
