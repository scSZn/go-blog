package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
	"github.com/scSZn/blog/pkg/app"
)

type ListTagParams struct {
	TagName       string
	ContainDelete bool
	Status        *uint8
	OrderKey      string
	Order         string
}

type TagDAO struct {
	db *gorm.DB
}

func NewTagDAO(ctx context.Context, db *gorm.DB) *TagDAO {
	return &TagDAO{
		db: db.WithContext(ctx),
	}
}

func (d *TagDAO) CreateTag(tag *model.Tag) (int64, error) {
	db := d.db.Table(model.TagTableName).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(tag)

	if err := db.Error; err != nil {
		return 0, errors.Wrapf(err, "TagDAO.CreateTag: create tag fail, tag is %+v", tag)
	}
	return db.RowsAffected, nil
}

func (d *TagDAO) DeleteTag(tagID string) (int64, error) {
	db := d.db.Table(model.TagTableName).Where("tag_id = ?", tagID).Updates(map[string]interface{}{
		"is_del": consts.DelStatus,
	})

	if err := db.Error; err != nil {
		return 0, errors.Wrapf(err, "TagDAO.DeleteTag: delete tag fail, tagID is %+v", tagID)
	}
	return db.RowsAffected, nil
}

func (d *TagDAO) ListTag(params *ListTagParams, pager *app.Pager) ([]*model.Tag, error) {
	db := d.db.Table(model.TagTableName)
	if params.TagName != "" {
		db = db.Where("tag_name like ?", params.TagName)
	}
	if params.Status != nil {
		db = db.Where("status = ?", *(params.Status))
	}
	if !params.ContainDelete {
		db = db.Where("is_del = ?", consts.NoDelStatus)
	}
	if params.OrderKey != "" && params.Order != "" {
		db = db.Order(fmt.Sprintf("%s %s", params.OrderKey, params.Order))
	}
	if pager != nil {
		db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())
	}
	var result []*model.Tag
	if err := db.Scan(&result).Error; err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrapf(err, "TagDAO.ListTag: list tag fail, params is %+v, pager is %+v", params, pager)
	}

	return result, nil
}

func (d *TagDAO) CountTag(params *ListTagParams) (int64, error) {
	db := d.db.Table(model.TagTableName)
	if params.TagName != "" {
		db = db.Where("tag_name like ?", params.TagName)
	}
	if params.Status != nil {
		db = db.Where("status = ?", *(params.Status))
	}
	if !params.ContainDelete {
		db = db.Where("is_del = ?", consts.NoDelStatus)
	}
	var result int64
	if err := db.Count(&result).Error; err != nil && err != sql.ErrNoRows {
		return 0, errors.Wrapf(err, "TagDAO.CountTag: count tag fail, params is %+v", params)
	}

	return result, nil
}

func (d *TagDAO) AddCount(tagIDs []string) error {
	err := d.db.Exec("UPDATE tag SET article_count = article_count + 1 WHERE tag_id in ?", tagIDs).Error
	if err != nil {
		return errors.Wrapf(err, "TagDAO.AddCount: update count fail, tagIDs is %+v", tagIDs)
	}
	return nil
}

// GetTagByTagIDBatch ??????????????????
func (d *TagDAO) GetTagByTagIDBatch(tagIDs []string) ([]*model.Tag, error) {
	var result []*model.Tag
	db := d.db.Table(model.TagTableName).Where("tag_id in ? AND is_del = ? AND status = ?", tagIDs, consts.NoDelStatus, consts.StatusEnable).Scan(&result)
	if err := db.Error; err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrapf(err, "TagDAO.GetTagByTagIDBatch: get tag fail, tagIDs is %+v", tagIDs)
	}
	return result, nil
}

func (d *TagDAO) UpdateTag(tagID string, status *uint8, newTagName string) (int64, error) {
	params := map[string]interface{}{
		"is_del": 0,
	}
	if status != nil {
		params["status"] = *status
	}
	newTagName = strings.TrimSpace(newTagName)
	if newTagName != "" {
		params["tag_name"] = newTagName
	}
	db := d.db.Table(model.TagTableName).Where("tag_id = ?", tagID).Updates(params)
	if err := db.Error; err != nil {
		return 0, errors.Wrapf(err, "TagDAO.UpdateTag: update tag fail, tagID is %+v, status is %+v, newTagName is %+v", tagID, status, newTagName)
	}
	return db.RowsAffected, nil
}
