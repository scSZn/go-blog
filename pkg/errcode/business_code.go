package errcode

var (
	LoginFail = NewError(1001, "login fail, username or password error")

	CreateArticleError = NewError(2001, "create article fail")
	ListArticleError   = NewError(2002, "list article fail")

	CreateTagError       = NewError(3001, "create tag fail")
	ListTagError         = NewError(3002, "list tag fail")
	TagAlreadyExistError = NewError(3003, "tag already exist")
	DeleteTagError       = NewError(3004, "删除标签失败")
	UpdateTagStatusError = NewError(3005, "更新标签状态失败")
)
