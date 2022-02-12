package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/internal/middleware"
	v1 "github.com/scSZn/blog/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Default())
	group := router.Group("/api/v1")
	{
		group.POST("/article", v1.CreateArticle)
	}
	return router
}
