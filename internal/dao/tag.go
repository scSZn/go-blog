package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
	"github.com/scSZn/blog/pkg/app"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ListTagParams struct {
	TagName  string
	IsDel    *bool
	Status   *uint8
	OrderKey string
	Order    string
}

type TagDAO struct {
	db *gorm.DB
}

func NewTagDAO(db *gorm.DB) *TagDAO {
	return &TagDAO{
		db: db,
	}
}

func (d *TagDAO) CreateTag(tag *model.Tag) (int64, error) {
	db := d.db.Table(model.TagTableName).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(tag)

	if err := db.Error; err != nil {
		return 0, errors.Wrap(err, "TagDAO.CreateTag: create tag fail: ")
	}
	return db.RowsAffected, nil
}

func (d *TagDAO) ListTag(params *ListTagParams, pager *app.Pager) ([]*model.Tag, error) {
	db := d.db.Table(model.TagTableName)
	if params.TagName != "" {
		db = db.Where("tag_name like %?%", params.TagName)
	}
	if params.Status != nil {
		db = db.Where("status = ?", *(params.Status))
	}
	if params.IsDel != nil {
		db = db.Where("is_del = ?", *(params.IsDel))
	}
	if params.OrderKey != "" && params.Order != "" {
		db = db.Order(fmt.Sprintf("%s %s", params.OrderKey, params.Order))
	}
	if pager != nil {
		db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())
	}
	var result []*model.Tag
	if err := db.Scan(&result).Error; err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "TagDAO.ListTag: list tag fail: ")
	}

	return result, nil
}

func (d *TagDAO) CountTag(params *ListTagParams) (int64, error) {
	db := d.db.Table(model.TagTableName)
	if params.TagName != "" {
		db = db.Where("tag_name like %?%", params.TagName)
	}
	if params.Status != nil {
		db = db.Where("status = ?", *(params.Status))
	}
	if params.IsDel != nil {
		db = db.Where("is_del = ?", *(params.IsDel))
	}
	var result int64
	if err := db.Count(&result).Error; err != nil && err != sql.ErrNoRows {
		return 0, errors.Wrap(err, "TagDAO.CountTag: count tag fail: ")
	}

	return result, nil
}

func (d *TagDAO) AddCount(tagIDs []string) error {
	err := d.db.Exec("UPDATE tag SET article_count = article_count + 1 WHERE tag_id in ?", tagIDs).Error
	if err != nil {
		return errors.Wrap(err, "TagDAO.AddCount: update count fail: ")
	}
	return nil
}

// GetTagByTagIDBatch 批量获取标签
func (d *TagDAO) GetTagByTagIDBatch(tagID ...string) ([]*model.Tag, error) {
	var result []*model.Tag
	db := d.db.Table(model.TagTableName).Where("tag_id in ? AND is_del = ? AND status = ?", tagID, consts.NoDelStatus, consts.StatusEnable).Scan(&result)
	if err := db.Error; err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "TagDAO.GetTagByTagIDBatch: get tag fail")
	}
	return result, nil
}
