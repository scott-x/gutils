drop database if exists mydb;
create database mydb DEFAULT CHARACTER SET 'UTF8';
use mydb;
-- drop table if exists user;
create table user(
  id int auto_increment key comment '用户编号',
  username varchar(50) not null comment '用户名',
  password varchar(50) not null comment '密码',
  nickname varchar(50) comment '昵称',
  age int comment '年龄',
  sex varchar(5) comment '性别',
  phone varchar(13) comment '电话',
  email varchar(20) comment '邮箱',
  createTime datetime comment '用户创建时间',
  updateTime datetime comment '用户上次更新时间',
  lastLogin datetime comment '用户上次登陆时间',
  userStatus int comment '用户账号状态 0 正常 ，1 锁定 ， 2 删除',
  remark text comment '备注'
)ENGINE=INNODB CHARSET=UTF8;

select * from user;
insert into 
user(username,password,nickname,age,sex,phone,email,createTime,updateTime,lastLogin,userStatus,remark) values
	 ('zhangsan','zhangsan','章三',18,'男','18320987665','zhangsan@22k.life','2019-03-09','2019-03-09','2019-03-09',0,'暂无'),
	 ('lisi','lisi','礼四',26,'女','18320987666','lisi@22k.life','2019-03-09','2019-03-09','2019-03-09',1,'暂无'),
	 ('wangwu','wangwu','汪五',32,'男','18320987667','wang5@22k.life','2019-03-09','2019-03-09','2019-03-09',2,'暂无'),
	 ('admin','admin','super admin',22,'男','18320987668','admin@22k.life','2019-03-09','2019-03-09','2019-03-09',2,'暂无');

