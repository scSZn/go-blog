package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/dao"
	"github.com/scSZn/blog/internal/model"
	"github.com/scSZn/blog/pkg/app"
	"github.com/scSZn/blog/pkg/errcode"
)

type CreateTagRequest struct {
	TagName string `json:"tag_name"`
	app.Pager
}

type DeleteTagRequest struct {
	TagID string `json:"tag_id" uri:"tag_id"`
}

type ListTagRequest struct {
	TagName  string `json:"tag_name" form:"tag_name"`
	OrderKey string `json:"order_key" form:"order_key"`
	Order    string `json:"order" form:"order"`
	Status   *uint8 `json:"status" form:"status"`
	IsDel    *bool  `json:"-"`
	app.Pager
}

type UpdateTagStatusRequest struct {
	TagID  string `json:"tag_id" uri:"tag_id"`
	Status uint8  `json:"status"`
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
		global.Logger.Errorf(ts.ctx, "TagService.CreateTag: create tag fail, request is %v, err: %v", request, err)
		return errcode.CreateTagError
	}

	if rowAffected == 0 {
		global.Logger.Warnf(ts.ctx, "TagService.CreateTag: tag %s already exists", request.TagName)
		return errcode.TagAlreadyExistError
	}
	return nil
}

func (ts *TagService) DeleteTag(request *DeleteTagRequest) error {
	tagDao := dao.NewTagDAO(ts.db)

	_, err := tagDao.DeleteTag(request.TagID)
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.DeleteTag: delete tag fail, tag_id is %v, err: %v", request.TagID, err)
		return errcode.DeleteTagError
	}

	return nil
}

func (ts *TagService) ListTag(request *ListTagRequest) ([]*model.Tag, error) {
	tagDao := dao.NewTagDAO(ts.db)
	params := dao.ListTagParams{
		TagName:  fmt.Sprintf("%%%s%%", request.TagName),
		OrderKey: request.OrderKey,
		Order:    request.Order,
		Status:   request.Status,
		IsDel:    request.IsDel,
	}

	result, err := tagDao.ListTag(&params, &request.Pager)
	for _, tag := range result {
		if tag.IsDel {
			tag.Status = consts.StatusDeleted
		}
	}
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.ListTag: list tag fail, params is %#v, err: %v", params, err)
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
		global.Logger.Errorf(ts.ctx, "TagService.CountTag: count tag fail, params is %v, err: %v", params, err)
		return 0, errcode.ListTagError
	}

	return result, nil
}

func (ts *TagService) UpdateTagStatus(request *UpdateTagStatusRequest) error {
	tagDao := dao.NewTagDAO(ts.db)

	result, err := tagDao.UpdateTagStatus(request.TagID, request.Status)
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.UpdateTagStatus: update tag status fail, request is %v, err: %v", request, err)
		return errcode.UpdateTagStatusError
	}

	if result == 0 {
		global.Logger.Errorf(ts.ctx, "TagService.UpdateTagStatus: update tag status fail, unknown error, row affected is 0, request is %+v", request)
		return errcode.UpdateTagStatusError
	}

	return nil
}
