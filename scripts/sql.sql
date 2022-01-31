create database if not exists `blog_db`;
use `blog_db`;

drop table if exists `article`;
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

drop table if exists `tag_article`;
create table tag_article
(
    id         bigint auto_increment primary key,
    tag_id     varchar(64) null,
    article_id varchar(64) null,
    is_del     tinyint     null,
    constraint idx_tag_article unique (tag_id, article_id)
) charset utf8mb4;

drop table if exists `blog_ext`;
create table blog_ext
(
    id         bigint primary key auto_increment,
    article_id varchar(64) null,
    view_count int         null,
    like_count int         null,
    constraint blog_ext_ibfk_1 foreign key (article_id) references article (article_id)
) charset utf8mb4;

create index fk_article_ext on blog_ext (article_id);