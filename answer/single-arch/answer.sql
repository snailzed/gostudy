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