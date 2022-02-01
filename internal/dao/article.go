package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/model"
)

type ArticleDAO struct {
	DB *gorm.DB
}

func NewArticleDAO() *ArticleDAO {
	return &ArticleDAO{
		DB: global.DB,
	}
}

func (a *ArticleDAO) GetArticleByArticleID(ctx context.Context, articleID string) (*model.Article, error) {
	var result *model.Article
	err := a.DB.Table(model.ArticleTableName).Where("article_id = ?", articleID).First(&result).Error
	if err != nil {
		if err == sql.ErrNoRows {
			global.Logger.Infof(ctx, "[dao.GetArticleByArticleID] no rows, article_id: %s", articleID)
			return nil, nil
		}
		return nil, err
	}
	return result, err
}

func (a *ArticleDAO) CreateArticle(ctx context.Context, article *model.Article) error {
	var result *model.Article
	return a.DB.Model(result).Create(article).Error
}
