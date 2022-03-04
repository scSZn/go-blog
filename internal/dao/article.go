package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/scSZn/blog/consts"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/model"
	"github.com/scSZn/blog/pkg/app"
)

type ListArticleParams struct {
	TitleLike     string
	AuthorLike    string
	TagIDs        []string
	Status        *uint8
	ContainDelete bool
}

type ArticleDAO struct {
	db *gorm.DB
}

func NewArticleDAO(ctx context.Context, db *gorm.DB) *ArticleDAO {
	return &ArticleDAO{
		db: db.WithContext(ctx),
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
			global.Logger.Infof(ctx, map[string]interface{}{
				"params: ": fmt.Sprintf("articleID: %s", articleID),
			}, "[dao.GetArticleByArticleID] no rows")
			return nil, nil
		}
		return nil, err
	}
	return result, err
}

func (a *ArticleDAO) CreateArticle(article *model.Article) error {
	err := a.db.Create(article).Error
	if err != nil {
		return errors.Wrapf(err, "ArticleDAO.CreateArticle: create article fail, article: %+v", article)
	}
	return nil
}

func (a *ArticleDAO) UpdateArticle(article *model.Article) error {
	err := a.db.Updates(article).Where("article_id = ?", article.ArticleID).Error
	if err != nil {
		return errors.Wrapf(err, "ArticleDAO.UpdateArticle: update article fail, article: %+v", article)
	}
	return nil
}

// List 根据条件查询文章
// TODO: 是否接入ES
func (a *ArticleDAO) List(params *ListArticleParams, pager app.Pager) ([]*model.Article, error) {
	db := a.db.Table(model.ArticleTableName)
	params.TitleLike = strings.TrimSpace(params.TitleLike)
	params.AuthorLike = strings.TrimSpace(params.AuthorLike)

	// todo: like 空字符串是否有性能上的缺失，加了索引和不加索引分别进行测试
	if params.TitleLike != "" {
		db = db.Where("title like %?%", params.TitleLike)
	}
	if params.AuthorLike != "" {
		db = db.Where("author like %?%", params.AuthorLike)
	}
	if len(params.TagIDs) > 0 {
		db = db.Where("tag_id in %s", params.TagIDs)
	}
	if !params.ContainDelete && params.Status != nil && *params.Status != consts.StatusDeleted {
		db = db.Where("is_del = ?", consts.NoDelStatus)
	}
	if params.Status != nil {
		if *params.Status == consts.StatusDeleted {
			db = db.Where("is_del = ?", consts.DelStatus)
		} else {
			db = db.Where("status = ?", params.Status)
		}
	}

	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())
	var result []*model.Article
	if err := db.Scan(&result).Error; err != nil {
		return nil, errors.Wrap(err, "ArticleDAO.List: query article list fail")
	}
	return result, nil
}

// List 根据条件查询文章
// TODO: 是否接入ES
func (a *ArticleDAO) Count(params *ListArticleParams) (int64, error) {
	db := a.db.Table(model.ArticleTableName)
	params.TitleLike = strings.TrimSpace(params.TitleLike)
	params.AuthorLike = strings.TrimSpace(params.AuthorLike)

	// todo: like 空字符串是否有性能上的缺失，加了索引和不加索引分别进行测试
	if params.TitleLike != "" {
		db = db.Where("title like %?%", params.TitleLike)
	}
	if params.AuthorLike != "" {
		db = db.Where("author like %?%", params.AuthorLike)
	}
	if len(params.TagIDs) > 0 {
		db = db.Where("tag_id in %s", params.TagIDs)
	}
	if !params.ContainDelete && params.Status != nil && *params.Status != consts.StatusDeleted {
		db = db.Where("is_del = ?", consts.NoDelStatus)
	}
	if params.Status != nil {
		if *params.Status == consts.StatusDeleted {
			db = db.Where("is_del = ?", consts.DelStatus)
		} else {
			db = db.Where("status = ?", params.Status)
		}
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "ArticleDAO.Count: query article count fail")
	}
	return count, nil
}
