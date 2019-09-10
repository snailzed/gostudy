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