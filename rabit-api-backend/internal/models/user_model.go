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
	UserName       string            `gorm:"UserName"`
	UserAccount    string            `gorm:"UserName"`
	Method         string            `gorm:"method"`
	RequestHeader  map[string]string `gorm:"requestHeader"`
	ResponseHeader map[string]string `gorm:"responseHeader"`
	UserId         int               `gorm:"userId"`
	Status         int               `gorm:"status"`
	IsDelete       int               `gorm:"isDelete"`
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

func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}
