package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/scSZn/blog/internal/dao"
	"github.com/scSZn/blog/internal/model"
)

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

type ArticleService struct {
}

func NewArticleService() ArticleService {
	return ArticleService{}
}

func (as ArticleService) CreateArticle(ctx context.Context, request CreateArticleRequest) error {
	articleDao := dao.NewArticleDAO()
	article := &model.Article{
		Title:     request.Title,
		Author:    request.Author,
		Summary:   request.Summary,
		Content:   request.Content,
		ArticleID: uuid.New().String(),
	}
	return articleDao.CreateArticle(article)
}
