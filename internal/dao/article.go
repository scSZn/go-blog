package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/model"
)

type ArticleDAO struct {
	db *gorm.DB
}

func NewArticleDAO(db *gorm.DB) *ArticleDAO {
	return &ArticleDAO{
		db: db,
	}
}

func (a *ArticleDAO) GetDB() *gorm.DB {
	return a.db
}

func (a *ArticleDAO) GetArticleByArticleID(ctx context.Context, articleID string) (*model.Article, error) {
	var result *model.Article
	err := a.db.Table(model.ArticleTableName).Where("article_id = ?", articleID).First(&result).Error
	if err != nil {
		if err == sql.ErrNoRows {
			global.Logger.Infof(ctx, "[dao.GetArticleByArticleID] no rows, article_id: %s", articleID)
			return nil, nil
		}
		return nil, err
	}
	return result, err
}

func (a *ArticleDAO) CreateArticle(article *model.Article) error {
	return a.db.Create(article).Error
}
