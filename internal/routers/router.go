package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/internal/middleware"
	v1Admin "github.com/scSZn/blog/internal/routers/api/v1/admin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Default())
	apiV1 := router.Group("/api/v1")

	admin := apiV1.Group("/admin")
	{
		admin.POST("/articles", v1Admin.CreateArticle)
		admin.GET("/articles", v1Admin.ListArticleAdmin)

		admin.POST("/tags", v1Admin.CreateTag)
	}
	return router
}
