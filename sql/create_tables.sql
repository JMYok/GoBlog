create table blog_category
(
    cid       int auto_increment
        primary key,
    name      varchar(64)                        not null,
    create_at datetime default CURRENT_TIMESTAMP not null,
    update_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint idx_cid
        unique (cid)
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table blog_post
(
    pid         int auto_increment
        primary key,
    title       varchar(255)  not null,
    content     longtext      not null,
    markdown    longtext      not null,
    category_id int           not null,
    user_id     bigint        not null,
    view_count  int default 0 not null,
    type        int default 0 not null,
    slug        varchar(255)  null,
    create_at   datetime      not null,
    update_at   datetime      not null,
    constraint idx_pid
        unique (pid)
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table blog_user
(
    uid       bigint      not null
        primary key,
    user_name varchar(64) not null,
    passwd    varchar(64) not null,
    avatar    varchar(64) not null,
    create_at datetime    not null,
    update_at datetime    not null,
    constraint idx_uid
        unique (uid),
    constraint idx_user_name
        unique (user_name)
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table users
(
    id         bigint auto_increment
        primary key,
    openid     varchar(255) not null,
    email      varchar(255) not null,
    vip        tinyint      not null,
    sc         tinyint      not null,
    lt         tinyint      not null,
    start_time datetime     not null,
    end_time   datetime     not null,
    subscribe  varchar(255) not null,
    password   varchar(255) not null,
    phone      varchar(255) not null
)
    collate = utf8mb4_unicode_ci
    row_format = DYNAMIC;


