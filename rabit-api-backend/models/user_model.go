package models

import "gorm.io/gorm"

/*
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
*/

type UserModel struct {
	gorm.Model
	UserAccount  string  `json:"user_account" gorm:"size:256;unique;not null;comment:用户登录账号"` // 用户登录账号
	UserPassword string  `json:"user_password" gorm:"size:256;not null;comment:密码"`           // 用户的密码
	NickName     string  `json:"nick_name" gorm:"size:256;comment:用户昵称"`                      // 用户的昵称
	AvatarUrl    *string `json:"avatar_url,omitempty" gorm:"size:1024;comment:用户头像地址"`
	Gender       uint    `json:"gender" gorm:"comment:用户性别;size:2"`
	Phone        *string `json:"phone,omitempty" gorm:"size:256;comment:用户手机"`           // 用户头像
	Email        string  `json:"email" gorm:"size:256;comment:用户邮箱"`                     // 用户的邮箱
	RoleID       uint    `json:"role_id" gorm:"size:2;comment:用户的角色ID(0-普通用户; 1-Admin)"` // 用户的角色ID
	UserStatus   uint    `json:"user_status" gorm:"size:2;default:0;comment:用戶status(0-正常)"`
	AccessKey    string  `json:"access_key" gorm:"comment:用户认证;size:512"`
	SecretKey    string  `json:"secret_key" gorm:"comment:密钥;size:512"`
	// RoleModel    RoleModel `gorm:"foreignKey:RoleID;" json:"roleModel"`
}

type UserCreateRequest struct {
	UserAccount  string `json:"user_account" binding:"required"`
	UserPassword string `json:"user_password" bingding:"required"`
	NickName     string `json:"nick_name"` // user-nickname
	RoleID       uint   `json:"role_id"`
}

type UserInfo struct {
	ID          uint    `json:"id"`
	UserAccount string  `json:"user_account"`
	NickName    string  `json:"nick_name"`
	AvatarUrl   *string `json:"avatar_url"`
	Gender      uint    `json:"gender"`
	Phone       *string `json:"phone"`
	Email       string  `json:"email"`
	RoleID      uint    `json:"role_id"`
	UserStatus  uint    `json:"user_status"`
	AccessKey   string  `json:"access_key"`
	SecretKey   string  `json:"secret_key"`
}

type UserChangePassword struct {
	ID            uint   `json:"id" binding:"required"`
	UserPassword  string `json:"user_password" binding:"required,min=8"`
	CheckPassword string `json:"check_password" binding:"required,eqfield=UserPassword"`
}
