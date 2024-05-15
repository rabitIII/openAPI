DROP DATABASE IF EXISTS rabit_api;

CREATE DATABASE IF NOT EXISTS rabit_api;

# 用户信息
CREATE TABLE IF NOT EXISTS rabit_api. user
(
    id BIGINT AUTO_INCREMENT COMMENT '主键' PRIMARY KEY ,
    userName VARBINARY(256) NULL COMMENT '用户昵称',
    userAccount VARCHAR(256) NOT NULL COMMENT '用户账号',
    userPassword VARCHAR(512) NOT NULL COMMENT '账号密码',
    userAvatar VARCHAR(1024) NULL COMMENT '用户头像',
    gender TINYINT NULL COMMENT '性别',
    userRole INT DEFAULT 0 NOT NULL COMMENT '用户角色（0-user, 1-Admin）',
    `isDelete` TINYINT DEFAULT 0 NOT NULL COMMENT '删除状态（0-未删，1-已删）',
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    `updatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deletedAt` DATETIME DEFAULT NULL COMMENT '删除时间',
    CONSTRAINT uni_userAccount UNIQUE (userAccount)
) COMMENT '用户表';

# 接口信息
CREATE TABLE IF NOT EXISTS rabit_api. `interface_info`
(
    `id` INT AUTO_INCREMENT COMMENT '主键' PRIMARY KEY ,
    `name` VARBINARY(256) NOT NULL COMMENT '接口名称',
    `description` VARCHAR(256) NULL COMMENT '接口描述',
    `url` VARCHAR(512) NOT NULL COMMENT '接口地址',
    `method` VARCHAR(256) NULL COMMENT '接口类型',
    `requestHeader` TEXT NULL COMMENT '请求头',
    `responseHeader` TEXT NULL COMMENT '响应头',
    `userId` BIGINT NOT NULL COMMENT '创始人id',
    `status` INT DEFAULT 0 NOT NULL COMMENT '接口状态',
    `isDelete` TINYINT DEFAULT 0 NOT NULL COMMENT '删除状态（0-未删，1-已删）',
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    `updatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deletedAt` DATETIME DEFAULT NULL COMMENT '删除时间'
) COMMENT '接口信息';

insert into rabit_api.`interface_info` (`name`, `description`, `url`, `method`, `requestHeader`, `responseHeader`, `userId`, `status`) values ('许擎宇', '这是一个测试数据1', 'www.cary-king.net', 'GET', '请求头1', '响应头1', 9500534531, 0);
insert into rabit_api.`interface_info` (`name`, `description`, `url`, `method`, `requestHeader`, `responseHeader`, `userId`, `status`) values ('陆弘文', '这是一个测试数据2', 'www.leslee-kuhn.net', 'GET', '请求头2', '响应头2', 3982575846, 0);
insert into rabit_api.`interface_info` (`name`, `description`, `url`, `method`, `requestHeader`, `responseHeader`, `userId`, `status`) values ('毛建辉', '这是一个测试数据3', 'www.rosaria-kilback.io', 'POST', '请求头3', '响应头3', 121776355, 0);
insert into rabit_api.`interface_info` (`name`, `description`, `url`, `method`, `requestHeader`, `responseHeader`, `userId`, `status`) values ('彭雨泽', '这是一个测试数据4', 'www.norris-bergstrom.biz', 'PUT', '请求头4', '响应头4', 740, 0);
insert into rabit_api.`interface_info` (`name`, `description`, `url`, `method`, `requestHeader`, `responseHeader`, `userId`, `status`) values ('傅志强', '这是一个测试数据5', 'www.jordan-reinger.com', 'DELETE', '请求头5', '响应头5', 35542559, 0);
