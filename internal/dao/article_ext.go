package dao

import (
	"github.com/pkg/errors"
	"github.com/scSZn/blog/internal/model"
	"gorm.io/gorm"
)

type ArticleExtDAO struct {
	db *gorm.DB
}

func NewArticleExtDAO(db *gorm.DB) *ArticleExtDAO {
	return &ArticleExtDAO{
		db: db,
	}
}

// Create 创建扩展信息
func (d *ArticleExtDAO) Create(ext *model.ArticleExt) error {
	err := d.db.Table(model.ArticleExtTableName).Create(ext).Error
	if err != nil {
		return errors.Wrap(err, "ArticleExtDAO.Create: create article ext fail")
	}
	return nil
}
