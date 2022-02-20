package middleware

const tokenHeaderKey = "token"

//
//func TokenVerify() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		token := ctx.GetHeader(tokenHeaderKey)
//		path := ctx.Request.URL.Path
//		if path == "" {
//			path = ctx.Request.URL.RawPath
//		}
//		// 如果是登录界面
//		if strings.Contains(path, "login") {
//			ctx.Next()
//			return
//		}
//
//	}
//}
