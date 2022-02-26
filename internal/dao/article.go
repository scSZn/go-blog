package dao

import (
	"context"
	"database/sql"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/model"
	"github.com/scSZn/blog/pkg/app"
)

type ListArticleParams struct {
	TitleLike  string
	AuthorLike string
	TagIDs     []string
	Status     uint8
	IsDel      *bool
}

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
	if params.Status != 0 {
		db = db.Where("status = ?", params.Status)
	}
	if params.IsDel != nil {
		db = db.Where("is_del = ?", *(params.IsDel))
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
	if params.Status != 0 {
		db = db.Where("status = ?", params.Status)
	}
	if params.IsDel != nil {
		db = db.Where("is_del = ?", *(params.IsDel))
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "ArticleDAO.Count: query article count fail")
	}
	return count, nil
}
