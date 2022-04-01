package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
)

type TagArticleDAO struct {
	db *gorm.DB
}

func NewTagArticleDAO(ctx context.Context, db *gorm.DB) *TagArticleDAO {
	return &TagArticleDAO{
		db: db.WithContext(ctx),
	}
}

// CreateTagArticleBatch 为一个文章创建多个标签关联关系
func (d *TagArticleDAO) CreateTagArticleBatch(articleID string, tagIDs []string) error {
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
	if err := d.db.Create(tagArticles).Error; err != nil {
		return errors.Wrap(err, "TagArticleDAO.CreateTagArticleBatch: create tag_article relationship fail")
	}
	return nil
}

// GetTagIDsByArticleID 获取articleID文章的所有关联信息
func (d *TagArticleDAO) GetTagIDsByArticleID(articleID string) ([]*model.TagArticle, error) {
	var result []*model.TagArticle
	db := d.db.Table(model.TagArticleTableName).Where("article_id = ? AND is_del = ?", articleID, consts.NoDelStatus).Scan(&result)
	if err := db.Error; err != nil {
		return nil, errors.Wrap(err, "TagArticleDAO.GetTagIDsByArticleID: get article tag relationship fail")
	}
	return result, nil
}

// GetTagIDsByArticleIDBatch 批量获取articleID文章的所有关联信息
func (d *TagArticleDAO) GetTagIDsByArticleIDBatch(articleIDs []string) ([]*model.TagArticle, error) {
	var result []*model.TagArticle
	db := d.db.Table(model.TagArticleTableName).Where("article_id in ? AND is_del = ?", articleIDs, consts.NoDelStatus).Scan(&result)
	if err := db.Error; err != nil {
		return nil, errors.Wrap(err, "TagArticleDAO.GetTagIDsByArticleIDBatch: get article tag relationship fail")
	}
	return result, nil
}
