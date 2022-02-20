package errcode

var (
	Success = NewError(0, "success")

	BindError = NewError(400, "client request error")
)
