package models

type UserModel struct {
	Model
	UserAccount  string    `json:"-" gorm:"size:36;column:userAccount;unique;not null;comment:用户登录账号"` // 用户登录账号
	UserPassword string    `json:"-" gorm:"size:128;column:userPassword;comment:密码"`                   // 用户的密码
	NickName     string    `json:"nickName" gorm:"size:36;column:nickName;comment:用户昵称"`               // 用户的昵称
	Email        string    `json:"email" gorm:"size:128;column:email;comment:用户邮箱"`                    // 用户的邮箱
	Token        string    `json:"-" gorm:"size:36;column:token;comment:用户的唯一识别符"`                     // 其他平台的唯一id
	RoleID       uint      `json:"roleID" gorm:"size:36;column:roleID;comment:用户的角色ID（身份绑定）"`          // 用户的角色ID
	RoleModel    RoleModel `gorm:"foreignKey:RoleID;" json:"roleModel"`
}
