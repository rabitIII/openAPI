package models

import "time"

type LoginModel struct {
	UserId   uint      `json:"user_id"`
	NickName string    `json:"nick_name"`
	Token    string    `json:"token"`
	RoleID   uint      `json:"role_id"`
	LoginAt  time.Time `gorm:"autoUpdateTime" json:"login_time"`
}
