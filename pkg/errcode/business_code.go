package errcode

var (
	LoginFail = NewError(1001, "登录失败，用户名或密码错误")

	CreateArticleError = NewError(2001, "create article fail")
	ListArticleError   = NewError(2002, "list article fail")
	CountArticleError  = NewError(2003, "统计文章总数失败")

	CreateTagError       = NewError(3001, "create tag fail")
	ListTagError         = NewError(3002, "list tag fail")
	TagAlreadyExistError = NewError(3003, "标签名称已存在")
	DeleteTagError       = NewError(3004, "删除标签失败")
	UpdateTagError       = NewError(3005, "更新标签失败")
)
