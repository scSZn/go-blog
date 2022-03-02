package dto

import (
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
)

type ArticleBaseInfo struct {
	ArticleID     string       `json:"article_id" `
	Title         string       `json:"title" `
	Author        string       `json:"author" `
	Summary       string       `json:"summary" `
	BackgroundURL string       `json:"background_url" `
	Tags          []*model.Tag `json:"tags"`
	PublishTime   string       `json:"publish_time"`
	LikeCount     int64        `json:"like_count"`
	ViewCount     int64        `json:"view_count"`
	Status        uint8        `json:"status"`
}

func GenArticleBashInfoFromArticleModel(article *model.Article) *ArticleBaseInfo {
	return &ArticleBaseInfo{
		ArticleID:     article.ArticleID,
		Title:         article.Title,
		Author:        article.Author,
		Summary:       article.Summary,
		BackgroundURL: article.BackgroundURL,
		PublishTime:   article.CreatedAt.Format(consts.TimeFormatLayout),
		LikeCount:     article.LikeCount,
		ViewCount:     article.ViewCount,
		Status:        article.Status,
	}
}
