package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/dao"
	"github.com/scSZn/blog/internal/dto"
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
	TagName       string `json:"tag_name" form:"tag_name"`
	OrderKey      string `json:"order_key" form:"order_key"`
	Order         string `json:"order" form:"order"`
	Status        *uint8 `json:"status" form:"status"`
	ContainDelete bool   `json:"ContainDelete" form:"contain_delete"`
	app.Pager
}

type UpdateTagRequest struct {
	TagID   string `json:"tag_id" uri:"tag_id"`
	Status  *uint8 `json:"status"`
	TagName string `json:"tag_name"`
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
		global.Logger.Errorf(ts.ctx, "TagService.CreateTag: create tag fail, request is %v, err: %+v", request, err)
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
		global.Logger.Errorf(ts.ctx, "TagService.DeleteTag: delete tag fail, tag_id is %v, err: %+v", request.TagID, err)
		return errcode.DeleteTagError
	}

	return nil
}

func (ts *TagService) ListTag(request *ListTagRequest) ([]*dto.TagInfo, int64, error) {
	tagDao := dao.NewTagDAO(ts.db)

	params := dao.ListTagParams{
		TagName:       fmt.Sprintf("%%%s%%", request.TagName),
		OrderKey:      request.OrderKey,
		Order:         request.Order,
		Status:        request.Status,
		ContainDelete: request.ContainDelete, // 是否包含删除的数据
	}

	tags, err := tagDao.ListTag(&params, &request.Pager)
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.ListTag: list tag fail, params is %+v, err: %+v", params, err)
		return nil, 0, errcode.ListTagError
	}
	for _, tag := range tags {
		if tag.IsDel {
			tag.Status = consts.StatusDeleted
		}
	}
	var result = dto.GenTagInfoFromModelTag(tags)

	total, err := tagDao.CountTag(&params)
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.ListTag: count tag fail, params is %+v, err: %+v", params, err)
		return nil, 0, errcode.ListTagError
	}

	return result, total, nil
}

func (ts *TagService) CountTag(request *ListTagRequest) (int64, error) {
	tagDao := dao.NewTagDAO(ts.db)
	params := dao.ListTagParams{
		TagName:       request.TagName,
		OrderKey:      request.OrderKey,
		Order:         request.Order,
		Status:        request.Status,
		ContainDelete: request.ContainDelete,
	}

	result, err := tagDao.CountTag(&params)
	if err != nil {
		global.Logger.Errorf(ts.ctx, "TagService.CountTag: count tag fail, params is %+v, err: %+v", params, err)
		return 0, errcode.ListTagError
	}

	return result, nil
}

func (ts *TagService) UpdateTag(request *UpdateTagRequest) error {
	tagDao := dao.NewTagDAO(ts.db)

	// 如果没有传入status，则默认修改为开启
	var status uint8 = consts.StatusEnable
	if request.Status == nil {
		request.Status = &status
	}

	result, err := tagDao.UpdateTag(request.TagID, request.Status, request.TagName)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate entry") {
			global.Logger.Warnf(ts.ctx, "TagService.UpdateTag: update tag status fail, tag_name is already exists, request is %+v, err: %+v", request, err)
			return errcode.TagAlreadyExistError
		}
		global.Logger.Errorf(ts.ctx, "TagService.UpdateTag: update tag status fail, request is %+v, err: %+v", request, err)
		return errcode.UpdateTagError
	}

	if result == 0 {
		global.Logger.Errorf(ts.ctx, "TagService.UpdateTag: update tag status fail, unknown error, row affected is 0, request is %+v", request)
		return errcode.UpdateTagError
	}

	return nil
}
