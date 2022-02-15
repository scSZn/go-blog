package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/dao"
	"github.com/scSZn/blog/internal/dto"
	"github.com/scSZn/blog/internal/model"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/set"
)

type CreateArticleRequest struct {
	Title         string   `json:"title"`
	Author        string   `json:"author"`
	Summary       string   `json:"summary"`
	Content       string   `json:"content"`
	TagIDs        []string `json:"tag_ids"`
	BackgroundURL string   `json:"background_url"`
}

type ListArticleRequest struct {
	Title  string   `json:"title" form:"title"`
	TagIDs []string `json:"tag_ids" form:"tag_ids"`
	Author string   `json:"author" form:"author"` // todo: 采用多选的形式？？
	Status uint8    `json:"status" form:"status"` // 状态值，内部使用（1：草稿，2：发布，3：禁用，4：删除）
	IsDel  *bool    `json:"-"`
	app.Pager
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

func (as *ArticleService) CreateArticle(request *CreateArticleRequest) error {
	tx := as.db.Begin()
	articleDao := dao.NewArticleDAO(tx)
	tagArticleDao := dao.NewTagArticleDAO(tx)
	tagDao := dao.NewTagDAO(tx)
	articleExtDao := dao.NewArticleExtDAO(tx)

	// 创建文章
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

	// 创建文章扩展信息
	err = articleExtDao.Create(&model.ArticleExt{
		ArticleID: article.ArticleID,
		ViewCount: 0,
		LikeCount: 0,
	})
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ArticleService.CreateArticle: create article ext fail")
	}

	// 创建文章与标签的关联关系
	err = tagArticleDao.CreateTagArticleBatch(article.ArticleID, request.TagIDs...)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ArticleService.CreateArticle: create tag article relationship fail")
	}

	// 更新标签的文章数量
	err = tagDao.AddCount(request.TagIDs)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ArticleService.CreateArticle: update tag count fail")
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "ArticleService.CreateArticle: commit fail")
	}
	return nil
}

// List 获取文章列表
// request.Status 是否需要根据状态值来筛选文章
// request.IsDel 是否需要获取已被软删除的文章
func (as *ArticleService) List(request *ListArticleRequest) ([]*dto.ArticleBaseInfo, error) {

	articleDao := dao.NewArticleDAO(as.db)
	tagArticleDao := dao.NewTagArticleDAO(as.db)
	tagDao := dao.NewTagDAO(as.db)

	listParam := &dao.ListArticleParams{
		TitleLike:  request.Title,
		AuthorLike: request.Author,
		TagIDs:     request.TagIDs,
		Status:     request.Status,
		IsDel:      request.IsDel,
	}

	// 获取符合条件的文章
	articles, err := articleDao.List(listParam, request.Pager)
	if err != nil {
		return nil, errors.Wrap(err, "ArticleService.List: query articles fail")
	}

	// 组装文章ID
	articleIDs := make([]string, 0, len(articles))
	for _, article := range articles {
		articleIDs = append(articleIDs, article.ArticleID)
	}

	// 根据文章ID批量获取标签文章关联关系，并使用集合过滤
	tagArticles, err := tagArticleDao.GetTagIDsByArticleIDBatch(articleIDs...)
	tagIDSet := set.NewStringSet()
	for _, tagArticle := range tagArticles {
		tagIDSet.Add(tagArticle.TagID)
	}
	// 将标签文章关联信息封装成map[articleID][]tagID，方便查找
	var articleTagMap = make(map[string][]string, len(articles))
	for _, tagArticle := range tagArticles {
		articleID := tagArticle.ArticleID
		tagID := tagArticle.TagID
		if _, ok := articleTagMap[articleID]; ok {
			articleTagMap[articleID] = append(articleTagMap[articleID], tagID)
		} else {
			articleTagMap[articleID] = []string{tagID}
		}
	}

	// 批量获取标签信息
	tags, err := tagDao.GetTagByTagIDBatch(tagIDSet.Elements()...)
	if err != nil {
		return nil, errors.Wrap(err, "ArticleService.List: query tag fail")
	}

	// 将标签信息封装成map，方便查找
	var tagMap = make(map[string]*model.Tag, len(tags))
	for _, tag := range tags {
		tagMap[tag.TagID] = tag
	}

	var result = make([]*dto.ArticleBaseInfo, 0, len(articles))
	for _, article := range articles {
		articleBaseInfo := &dto.ArticleBaseInfo{
			Model:         article.Model,
			Title:         article.Title,
			Summary:       article.Summary,
			Author:        article.Author,
			ArticleID:     article.ArticleID,
			BackgroundURL: article.BackgroundURL,
		}
		tagIDs := articleTagMap[article.ArticleID]
		tags := make([]*model.Tag, 0, len(tagIDs))
		for _, tagID := range tagIDs {
			tags = append(tags, tagMap[tagID])
		}
		articleBaseInfo.Tags = tags
		result = append(result, articleBaseInfo)
	}

	return result, nil
}

// Count 获取符合条件的文章数量
// request.Status 是否需要根据状态值来筛选文章
// request.IsDel 是否需要获取已被软删除的文章
func (as *ArticleService) Count(request *ListArticleRequest) (int64, error) {

	articleDao := dao.NewArticleDAO(as.db)

	listParam := &dao.ListArticleParams{
		TitleLike:  request.Title,
		AuthorLike: request.Author,
		TagIDs:     request.TagIDs,
		Status:     request.Status,
	}

	// 获取符合条件的文章
	count, err := articleDao.Count(listParam)
	if err != nil {
		return count, errors.Wrap(err, "ArticleService.List: query articles fail")
	}
	return count, nil
}
