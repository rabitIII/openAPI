package models

// RoleModel 角色权限表: 1.系统管理员; 4.用户
type RoleModel struct {
	Model
	RoleName string `gorm:"size:16;not null;comment:角色名称" json:"roleName"`                         // 角色名
	Pwd      string `gorm:"size:64;comment:角色密码" json:"-"`                                         // 角色密码
	IsSystem bool   `gorm:"column:isSystem;default:1;comment:是否为系统（0:用户,1:管理员）角色" json:"isSystem"` // 是否为系统（管理员）角色 1:System; 0:User
}
