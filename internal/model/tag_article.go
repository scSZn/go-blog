package model

/*
drop table if exists `tag_article`;
create table tag_article
(
    id         bigint auto_increment primary key,
    tag_id     varchar(64) null,
    article_id varchar(64) null,
    is_del     tinyint     null,
    constraint idx_tag_article unique (tag_id, article_id)
) charset utf8mb4;
*/

const TagArticleTableName = "tag_article"

// TagArticle 文章标签关联表
type TagArticle struct {
	ID        uint   `gorm:"primaryKey"`
	TagID     string `json:"tag_id" gorm:"column:tag_id"`
	ArticleID string `json:"article_id" gorm:"column:article_id"`
	IsDel     bool   `json:"is_del" gorm:"column:is_del"`
}

func (t *TagArticle) TableName() string {
	return TagArticleTableName
}
