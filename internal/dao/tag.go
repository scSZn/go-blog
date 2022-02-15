package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
	"gorm.io/gorm"
)

type TagDAO struct {
	db *gorm.DB
}

func NewTagDAO(db *gorm.DB) *TagDAO {
	return &TagDAO{
		db: db,
	}
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
