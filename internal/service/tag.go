package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/dao"
	"github.com/scSZn/blog/internal/model"
	"gorm.io/gorm"
)

type CreateTagRequest struct {
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
		TagName:      request.TagName,
		TagID:        uuid.New().String(),
		ArticleCount: 0,
	}

	err := tagDao.CreateTag(tag)
	if err != nil {
		return errors.Wrap(err, "TagService.Create: create tag fail: ")
	}

	return nil
}
