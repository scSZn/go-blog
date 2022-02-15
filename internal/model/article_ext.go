package model

const ArticleExtTableName = "article_ext"

type ArticleExt struct {
	ArticleID string `json:"article_id" gorm:"column:article_id;notnull"`
	ViewCount int64  `json:"view_count" gorm:"column:view_count"`
	LikeCount int64  `json:"like_count" gorm:"column:like_count"`
}
