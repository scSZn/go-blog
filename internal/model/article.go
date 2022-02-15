package model

/*
create table article
(
    id         bigint auto_increment primary key,
    article_id varchar(64)                not null,
    title      varchar(1024)              null,
    author     varchar(128)               null,
    summary    varchar(2048) charset utf8 null,
    content    longtext                   null,
    status     tinyint                    null,
    create_at  datetime                   null,
    update_at  datetime                   null,
    is_del     tinyint                    null,
    delete_at  datetime                   null,
    constraint article_id unique (article_id)
) charset utf8mb4;
*/

const ArticleTableName = "article"

// Article 文章结构体
type Article struct {
	*Model
	ArticleID     string `json:"article_id" gorm:"column:article_id;unique"`
	Title         string `json:"title" gorm:"column:title"`
	Author        string `json:"author" gorm:"column:author"`
	Summary       string `json:"summary" gorm:"column:summary"`
	BackgroundURL string `json:"background_url" gorm:"column:background_url"`
	Content       string `json:"content" gorm:"column:content"`
	ArticleExt    `gorm:"-"`
}

func (a *Article) TableName() string {
	return ArticleTableName
}
