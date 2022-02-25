create database if not exists `blog_db`;
use `blog_db`;

drop table if exists `article`;
create table article
(
    id             bigint auto_increment primary key,
    article_id     varchar(64)   not null,
    title          varchar(1024) not null,
    author         varchar(128)  not null,
    summary        varchar(2048) null,
    background_url varchar(1024) null,
    content        longtext      not null,
    view_count     int           not null,
    like_count     int           not null,
    status         tinyint       not null,
    created_at     datetime      null,
    updated_at     datetime      null,
    is_del         tinyint       not null,
    deleted_at     datetime      null,
    unique key uni_article_id (article_id)
) charset utf8mb4;

drop table if exists `tag`;
create table tag
(
    id            bigint auto_increment primary key,
    tag_id        varchar(64) not null,
    tag_name      varchar(20) not null,
    article_count int         not null,
    status        tinyint     not null,
    created_at    datetime    null,
    updated_at    datetime    null,
    deleted_at    datetime    null,
    is_del        tinyint     not null,
    unique key uni_tag_id (tag_id),
    unique key uni_tag_name (tag_name)
) charset utf8mb4;

drop table if exists `tag_article`;
create table tag_article
(
    id         bigint auto_increment primary key,
    tag_id     varchar(64) not null,
    article_id varchar(64) not null,
    is_del     tinyint     not null,
    unique key uni_tag_article (tag_id, article_id),
    key idx_tag_id (tag_id),
    key idx_article_id (article_id)
) charset utf8mb4;

drop table if exists user;
create table user
(
    id         bigint auto_increment primary key,
    uid        varchar(64)  not null,
    username   varchar(100) not null,
    passport   varchar(256) not null,
    nickname   varchar(100) not null,
    role       varchar(100) not null,
    avatar varchar(1000) not null,
    status     tinyint      not null,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    is_del     tinyint      not null,
    unique key uni_uid (uid),
    unique key uni_username (username)
) charset utf8mb4;

insert into user(`uid`, `username`, `passport`, `nickname`, `role`, `avatar`, `status`, `created_at`, `updated_at`,
                    `is_del`)
values ('1', 'admin',
        '243261243130247a3539424f46356d6751514b68783174324658723075414c6677467a542f666d4f493552474762344e356e5a66326c677a346d3661',
        '', 'admin', '', 20, now(), now(), 0)