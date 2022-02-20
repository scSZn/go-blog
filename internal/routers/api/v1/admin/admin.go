package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/service"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
)

func Login(ctx *gin.Context) {
	request := &service.LoginRequest{}
	response := app.NewResponse(ctx)
	err := ctx.Bind(request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.Login: bind error: %v", err)
		response.ReturnError(errcode.BindError)
		return
	}

	svc := service.NewLoginService(ctx)
	err = svc.Login(request)
	if err != nil {
		global.Logger.Errorf(ctx, "admin.Login: login fail: %v", err)
		response.ReturnError(errcode.LoginFail)
		return
	}

	response.ReturnData("12345678")
}

func Info(ctx *gin.Context) {
	//request := &service.LoginRequest{}
	response := app.NewResponse(ctx)
	//err := ctx.Bind(request)
	//if err != nil {
	//	global.Logger.Errorf(ctx, "admin.Login: bind error: %v", err)
	//	response.ReturnError(errcode.BindError)
	//	return
	//}
	//
	//svc := service.NewLoginService(ctx)
	//err = svc.Login(request)
	//if err != nil {
	//	global.Logger.Errorf(ctx, "admin.Login: login fail: %v", err)
	//	response.ReturnError(errcode.LoginFail)
	//	return
	//}

	//ctx.Header(consts.TokenHeaderKey, "12345678")
	response.ReturnData(gin.H{
		"roles":        []string{"admin"},
		"introduction": "I am a super administrator",
		"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"name":         "Super Admin",
	})
}
