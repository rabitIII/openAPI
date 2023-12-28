package models

type UserModel struct {
	Model
	UserName  string    `json:"-" gorm:"size:36;column:userName;unique;not null;comment:用户名"` // 用户名
	Password  string    `json:"-" gorm:"size:128;column:password;comment:密码"`                 // 用户的密码
	NickName  string    `json:"nickName" gorm:"size:36;column:nickName;comment:用户昵称"`         // 用户的昵称
	Email     string    `json:"email" gorm:"size:128;column:email;comment:用户邮箱"`              // 用户的邮箱
	Token     string    `json:"-" gorm:"size:36;column:token;comment:用户的唯一识别符"`               // 其他平台的唯一id
	RoleID    uint      `json:"roleID" gorm:"size:36;column:roleID;comment:用户对应的权限级别"`        // 用户对应的角色（拥有的权限）
	RoleModel RoleModel `gorm:"foreignKey:RoleID" json:"-"`
}
