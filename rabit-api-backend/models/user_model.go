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
	UserAccount  string `json:"-" gorm:"size:36;column:userAccount;unique;not null;comment:用户登录账号"`   // 用户登录账号
	UserPassword string `json:"-" gorm:"size:128;column:userPassword;comment:密码"`                     // 用户的密码
	NickName     string `json:"nickName" gorm:"size:36;column:nickName;comment:用户昵称"`                 // 用户的昵称
	Email        string `json:"email" gorm:"size:128;column:email;comment:用户邮箱"`                      // 用户的邮箱
	Token        string `json:"-" gorm:"size:36;column:token;comment:用户的唯一识别符"`                       // 其他平台的唯一id
	RoleID       uint   `json:"roleID" gorm:"size:36;column:roleID;comment:用户的角色ID(0-普通用户; 1-Admin)"` // 用户的角色ID
	AccessKey    string `json:"accessKey" gorm:"comment:用户认证;size:512"`
	SecretKey    string `json:"secretKey" gorm:"comment:密钥;size:512"`
	// RoleModel    RoleModel `gorm:"foreignKey:RoleID;" json:"roleModel"`
}

type UserCreateRequest struct {
	UserAccount  string `json:"userAccount" binding:"required"`
	UserPassword string `json:"userPassword" bingding:"required"`
	NickName     string `json:"nickName"` // user-nickname
	RoleID       uint   `json:"roleId" bingding:"required"`
}

// Create	创建用户
func (u *UserModel) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

// FindUserByID	查询用户
func FindUserByID(db *gorm.DB, id uint) (*UserModel, error) {
	var user UserModel
	res := db.First(&user, id)
	return &user, res.Error
}

func (u *UserModel) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *UserModel) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}
