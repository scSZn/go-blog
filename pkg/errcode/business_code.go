package errcode

var (
	CreateArticleError = NewError(1001, "create article fail")
	ListArticleError   = NewError(1002, "list article fail")

	CreateTagError       = NewError(2001, "create tag fail")
	TagAlreadyExistError = NewError(2101, "tag already exist")
)
