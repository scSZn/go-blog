package model

/*
drop table if exists `tag`;
create table tag
(
    id            bigint auto_increment primary key,
    tag_id        varchar(64) not null,
    tag_name      varchar(20) null,
    article_count int         null,
    status        tinyint     null,
    create_at     datetime    null,
    update_at     datetime    null,
    delete_at     datetime    null,
    is_del        tinyint     null,
    constraint tag_id unique (tag_id)
) charset utf8mb4;
*/

const TagTableName = "tag"

// Tag 标签结构体
type Tag struct {
	*Model
	TagID        string `json:"tag_id" gorm:"column:tag_id,unique"`
	TagName      string `json:"tag_name" gorm:"column:tag_name"`
	ArticleCount uint32 `json:"article_count" gorm:"column:article_count"`
}

func (t *Tag) TableName() string {
	return TagTableName
}
