package models

import "time"

type LoginModel struct {
	UserID    uint      `gorm:"colum:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	NickName  string    `gorm:"column:nickName;size:42" json:"NickName"`
	Token     string    `gorm:"size:256" json:"token"`
	RoleID    uint      `json:"role" binding:"required"` // 角色id
	LoginAt   time.Time `gorm:"autoUpdateTime" json:"createAt"`
}
