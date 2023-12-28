package models

type RoleModel struct {
	Model
	Title    string `gorm:"size:16;not null;comment:角色名称" json:"title"`                            // 角色名
	Pwd      string `gorm:"size:64;comment:角色密码" json:"-"`                                         // 角色密码
	IsSystem bool   `gorm:"column:isSystem;default:1;comment:是否为系统（1:管理员,0:用户）角色" json:"isSystem"` // 是否为系统（管理员）角色 1:System; 0:User
}
