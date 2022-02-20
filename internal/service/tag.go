package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/dao"
	"github.com/scSZn/blog/internal/model"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
	"gorm.io/gorm"
)

type CreateTagRequest struct {
	TagName string `json:"tag_name"`
	app.Pager
}

type ListTagRequest struct {
	TagName  string `json:"tag_name"`
	OrderKey string `json:"order_key"`
	Order    string `json:"order"`
	Status   *uint8 `json:"status"`
	IsDel    *bool  `json:"-"`
	app.Pager
}

type TagService struct {
	ctx context.Context
	db  *gorm.DB
}

func NewTagService(ctx context.Context) *TagService {
	return &TagService{
		ctx: ctx,
		db:  global.DB,
	}
}

func (ts *TagService) CreateTag(request *CreateTagRequest) error {
	tagDao := dao.NewTagDAO(ts.db)
	tag := &model.Tag{
		Model: &model.Model{
			Status: 20,
		},
		TagName:      request.TagName,
		TagID:        uuid.New().String(),
		ArticleCount: 0,
	}

	rowAffected, err := tagDao.CreateTag(tag)
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.Create: create tag fail: ", err)
		return errcode.CreateTagError
	}

	if rowAffected == 0 {
		global.Logger.Warnf(ts.ctx, "TagService.Create: tag %s already exists", request.TagName)
		return errcode.TagAlreadyExistError
	}
	return nil
}

func (ts *TagService) ListTag(request *ListTagRequest) ([]*model.Tag, error) {
	tagDao := dao.NewTagDAO(ts.db)
	params := dao.ListTagParams{
		TagName:  request.TagName,
		OrderKey: request.OrderKey,
		Order:    request.Order,
		Status:   request.Status,
		IsDel:    request.IsDel,
	}

	result, err := tagDao.ListTag(&params, &request.Pager)
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.ListTag: list tag fail: %v", err)
		return nil, errcode.ListTagError
	}

	return result, nil
}

func (ts *TagService) CountTag(request *ListTagRequest) (int64, error) {
	tagDao := dao.NewTagDAO(ts.db)
	params := dao.ListTagParams{
		TagName:  request.TagName,
		OrderKey: request.OrderKey,
		Order:    request.Order,
		Status:   request.Status,
		IsDel:    request.IsDel,
	}

	result, err := tagDao.CountTag(&params)
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.CountTag: count tag fail: %v", err)
		return 0, errcode.ListTagError
	}

	return result, nil
}
