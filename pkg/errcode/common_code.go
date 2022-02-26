package errcode

var (
	Success = NewError(0, "success")

	ClientRequestError         = NewError(400, "客户端请求错误")
	AuthorizationTokenTimeout  = NewError(40301, "token已过期")
	AuthorizationTokenInvalid  = NewError(40302, "token不可用")
	AuthorizationNotPermission = NewError(40303, "无权进行此项操作")

	ServerError = NewError(500, "服务端内部错误")
)
