package dto

import "github.com/scSZn/blog/internal/model"

type ArticleBaseInfo struct {
	*model.Model
	ArticleID     string       `json:"article_id" `
	Title         string       `json:"title" `
	Author        string       `json:"author" `
	Summary       string       `json:"summary" `
	BackgroundURL string       `json:"background_url" `
	Tags          []*model.Tag `json:"tags"`
}
