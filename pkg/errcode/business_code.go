package errcode

var (
	LoginFail = NewError(1001, "login fail, username or password error")

	CreateArticleError = NewError(2001, "create article fail")
	ListArticleError   = NewError(2002, "list article fail")

	CreateTagError       = NewError(3001, "create tag fail")
	TagAlreadyExistError = NewError(3101, "tag already exist")
)
