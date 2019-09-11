CREATE DATABASE mercury character set utf8mb4;
use mercury;

CREATE TABLE user (
  id bigint(20) primary key auto_increment not null ,
  username varchar(64) not null DEFAULT '' comment 'username',
  nickname varchar(64) NOT NULL DEFAULT '' COMMENT 'nickname',
  email varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  password varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  salt varchar(32) NOT NULL DEFAULT '' COMMENT '盐加密',
  sex tinyint(3) DEFAULT 0 COMMENT '0：未知，1男2女',
  create_time int default 0 comment '',
  update_time int default 0 comment ''
);

CREATE TABLE category (
  id bigint(20) primary key auto_increment not null ,
  name varchar(32) not null unique comment '分类名',
  create_time int default 0 comment '创建时间',
  update_time int default 0 comment '更新时间'
);

CREATE TABLE question (
  id bigint(20) primary key auto_increment not null ,
  caption varchar(256) not null default '' comment '标题',
  content varchar(4096) not null default '' comment '内容',
  author_id  bigint(20) not null default 0 comment '作者ID',
  category_id  bigint(20) not null default 0 comment '分类ID',
  status tinyint(3) default 0 comment '状态',
  create_time int default 0 comment '创建时间',
  update_time int default 0 comment '更新时间',
  index(status)
)engine=innodb character set utf8mb4;

CREATE TABLE answer(
    id bigint(20) primary key auto_increment not null ,
    question_id bigint(20)  not null default 0 comment '问题ID',
    content varchar(4096) not null default '' comment '回答内容',
    vote_count int default 0 comment '点赞数',
    like_count int default 0 comment '收藏数',
    create_time int default 0 comment '创建时间',
    update_time int default 0 comment '更新时间',
    index(question_id)
)engine=innodb character set utf8mb4;


-- 一个获取 root_comments接口，获取该答案下的所有顶级评论
-- 一个获取 child_comments接口，获取该答案下的所有子评论
-- 只支持两级评论
CREATE TABLE comment (
    id bigint(20) primary key auto_increment not null ,
    answer_id bigint(20)  not null default 0 comment '回答ID',
    author_id bigint(20) not null default 0 comment '用户ID',
    reply_author_id bigint(20) not null default 0 comment '被回复评论的用户ID',
    content varchar(1024) not null default '' comment '评论内容',
    parent_id bigint(20) not null default 0 comment '父级评论ID,0:顶级评论，其他否',
    vote_count int default 0 comment '点赞数',
    is_featured tinyint(3) default 0 comment '是否精选，0：否，1：是',
    is_author tinyint(3) default 0 comment '是否回答该问题的用户，0：否，1：是',
    child_comment_count int default 0 comment '子评论数',
    is_delete tinyint(3) default 0 comment '是否已删除，0：否，1：是',
    allow_vote tinyint(3) default 0 comment '是否允许点赞，0：是，1：否',
    allow_reply tinyint(3) default 0 comment '是否允许回复，0：是，1：否',
    create_time int default 0 comment '创建时间',
    update_time int default 0 comment '更新时间',
    index(answer_id),
    index(author_id),
    index(parent_id),
    index(is_featured),
    index(create_time)
)
