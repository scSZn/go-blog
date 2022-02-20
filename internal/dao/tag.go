package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
	"github.com/scSZn/blog/pkg/app"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagDAO struct {
	db *gorm.DB
}

func NewTagDAO(db *gorm.DB) *TagDAO {
	return &TagDAO{
		db: db,
	}
}

func (d *TagDAO) CreateTag(tag *model.Tag) error {
	db := d.db.Table(model.TagTableName).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(tag)

	if err := db.Error; err != nil {
		return errors.Wrap(err, "TagDAO.CreateTag: create tag fail: ")
	}
	if db.RowsAffected == 0 {
		return app.TagAlreadyExistError
	}
	return nil
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
