create database if not exists `blog_db`;
use `blog_db`;

drop table if exists `article`;
create table article
(
    id             bigint auto_increment primary key,
    article_id     varchar(64)   not null,
    title          varchar(1024) null,
    author         varchar(128)  null,
    summary        varchar(2048) null,
    background_url varchar(1024) null,
    content        longtext      null,
    status         tinyint       null,
    created_at     datetime      null,
    updated_at     datetime      null,
    is_del         tinyint       null,
    deleted_at     datetime      null,
    constraint uni_article_id unique (article_id)
) charset utf8mb4;

drop table if exists `tag`;
create table tag
(
    id            bigint auto_increment primary key,
    tag_id        varchar(64) not null,
    tag_name      varchar(20) null,
    article_count int         null,
    status        tinyint     null,
    created_at    datetime    null,
    updated_at    datetime    null,
    deleted_at    datetime    null,
    is_del        tinyint     null,
    constraint uni_tag_id unique (tag_id)
) charset utf8mb4;

drop table if exists `tag_article`;
create table tag_article
(
    id         bigint auto_increment primary key,
    tag_id     varchar(64) null,
    article_id varchar(64) null,
    is_del     tinyint     null,
    constraint uni_tag_article unique (tag_id, article_id)
) charset utf8mb4;

drop table if exists `article_ext`;
create table article_ext
(
    id         bigint primary key auto_increment,
    article_id varchar(64) null,
    view_count int         null,
    like_count int         null
) charset utf8mb4;

create index fk_article_ext on article_ext (article_id);