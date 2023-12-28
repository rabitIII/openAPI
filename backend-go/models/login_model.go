package models

type LoginModel struct {
	Model
	UserID    uint      `gorm:"colum:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	NickName  string    `gorm:"column:nickName;size:42" json:"NickName"`
	Token     string    `gorm:"size:256" json:"token"`
}
