package dao

import (
	"github.com/scSZn/blog/internal/model"
	"gorm.io/gorm"
)

type TagArticleDAO struct {
	db *gorm.DB
}

func NewTagArticleDAO(db *gorm.DB) *TagArticleDAO {
	return &TagArticleDAO{
		db: db,
	}
}

// CreateTagArticleBatch 为一个文章创建多个标签关联关系
func (d *TagArticleDAO) CreateTagArticleBatch(articleID string, tagIDs ...string) error {
	if len(tagIDs) == 0 {
		return nil
	}
	tagArticles := make([]*model.TagArticle, 0, len(tagIDs))
	for _, tagID := range tagIDs {
		tagArticles = append(tagArticles, &model.TagArticle{
			ArticleID: articleID,
			TagID:     tagID,
		})
	}
	return d.db.CreateInBatches(tagArticles, 10).Error
}
