package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/dao"
	"github.com/scSZn/blog/internal/model"
)

type CreateArticleRequest struct {
	Title         string   `json:"title"`
	Author        string   `json:"author"`
	Summary       string   `json:"summary"`
	Content       string   `json:"content"`
	TagIDs        []string `json:"tag_ids"`
	BackgroundURL string   `json:"background_url"`
}

type ArticleService struct {
	ctx context.Context
	db  *gorm.DB
}

func NewArticleService(ctx context.Context) *ArticleService {
	service := &ArticleService{
		ctx: ctx,
	}
	service.db = global.DB
	return service
}

func (as *ArticleService) CreateArticle(request CreateArticleRequest) error {
	tx := as.db.Begin()
	articleDao := dao.NewArticleDAO(tx)
	tagArticleDao := dao.NewTagArticleDAO(tx)
	article := &model.Article{
		Title:         request.Title,
		Author:        request.Author,
		Summary:       request.Summary,
		BackgroundURL: request.BackgroundURL,
		Content:       request.Content,
		ArticleID:     uuid.New().String(), // todo: ArticleID 使用分布式ID
	}

	err := articleDao.CreateArticle(article)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ArticleService.CreateArticle: create article fail")
	}

	err = tagArticleDao.CreateTagArticleBatch(article.ArticleID, request.TagIDs...)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ArticleService.CreateArticle: create tag article relationship fail")
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ArticleService.CreateArticle: commit fail")
	}
	return nil
}
