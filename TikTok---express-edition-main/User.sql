-- 创建库
create database if not exists ByteDance;

-- 切换库
use ByteDance;

-- 用户表
create table if not exists user
(
    id             bigint auto_increment comment 'id' primary key,
    userName       varchar(256)                       null comment '用户昵称',
    userAccount    varchar(256)                       not null comment '账号',
    userAvatar     varchar(1024)                      null comment '用户头像',
    gender         tinyint                            null comment '性别',
    userRole       int      default 0                 not null comment '用户角色 0 - 普通用户 1 - 管理员',
    userPassword   varchar(512)                       not null comment '密码',
    `accessKey`    varchar(512)                       null comment 'accessKey',
    `secretKey`    varchar(512)                       null comment 'secretKey',
    createTime     datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    follow_count   int      default 0                 null comment '关注总数',
    follower_count int      default 0                 null comment '粉丝总数',
    isDelete       tinyint  default 0                 not null comment '是否删除',
    is_favorite    tinyint  default 0                 not null comment '点赞状态',
    is_follow      tinyint  default 0                 not null comment '关注状态',
    constraint uni_userAccount
        unique (userAccount)
) comment '用户';
-- 视频表
create table if not exists video
(
    id         bigint comment 'id' primary key,
    userId     varchar(256)                       not null comment '作者Id',
    play_url   varchar(256)                       not null comment '视频地址',
    cover_url  varchar(256)                       not null comment '封面地址',
    title      varchar(256)                       not null comment '视频标题',
    isDelete   tinyint  default 0                 not null comment '是否删除',
    created_at datetime default CURRENT_TIMESTAMP not null comment '发布时间',
    playstatus tinyint  default 0                 not null comment '视频状态'-- 0 表示正常 1 表示不正常

) comment '视频';
-- 评论表
create table if not exists comment
(
    `id`            int(11)                                                                           NOT NULL AUTO_INCREMENT COMMENT '评论ID',
    `video_id`      bigint(20)                                                                        NOT NULL COMMENT '对应视频品ID',
    `content`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '评论内容',
    `created_date`    datetime(0)                                             DEFAULT current_timestamp not NULL COMMENT '创建时间',
    `action_type`    int(10)                                           DEFAULT current_timestamp not NULL COMMENT '操作类型',
    PRIMARY KEY (`id`) USING BTREE
) comment '评论';

-- 视屏点赞功能
create table if not exists ByteDance.`video_like`
(
    `id`       int                                not null auto_increment comment '主键id' primary key,
    `video_id` int                                not null comment '视屏id',
    `user_id`  int                                not null comment '用户id',
    `isLike`   tinyint  default 0                 not null comment '点赞状态 0-没点赞 1-点赞',
    createTime datetime default CURRENT_TIMESTAMP not null comment '创建时间'
) comment '视屏点赞功能';

insert into ByteDance.`video_like` (`video_id`, `user_id`) values (139669485, 75278);
insert into ByteDance.`video_like` (`video_id`, `user_id`) values (7368406, 886);
