package errcode

var (
	Success = NewError(0, "success")

	ClientRequestError        = NewError(400, "客户端请求错误")
	AuthorizationTokenTimeout = NewError(403, "token已过期")
	AuthorizationTokenInvalid = NewError(403, "token不可用")

	ServerError = NewError(500, "服务端内部错误")
)
