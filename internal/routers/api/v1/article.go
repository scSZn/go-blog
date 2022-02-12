package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/service"
)

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

func CreateArticle(ctx *gin.Context) {
	var request CreateArticleRequest
	err := ctx.Bind(&request)
	if err != nil {
		global.Logger.Errorf(ctx, "bind error, err: %v", err)
		return
	}

	svc := service.NewArticleService()
	svcRequest := service.CreateArticleRequest{
		Title:   request.Title,
		Summary: request.Summary,
		Content: request.Content,
		Author:  request.Author,
	}

	err = svc.CreateArticle(ctx, svcRequest)
	if err != nil {
		global.Logger.Errorf(ctx, "bind error, err: %v", err)
		ctx.JSON(500, "create article error")
		return
	}
	ctx.JSON(200, "create article success")
}
