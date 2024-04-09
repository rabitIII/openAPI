package models

import "gorm.io/gorm"

// InterfaceInfo api接口
type InterfaceInfo struct {
	gorm.Model
	UserId         uint   `json:"UserId" gorm:"comment:创建人的ID;not null"`
	Title          string `json:"title" gorm:"comment:接口名称;size:256;not null"`
	Description    string `json:"description" gorm:"comment:描述;size:256;"`
	Url            string `json:"url" gorm:"comment:接口地址；size:512;"`
	Method         string `json:"method" gorm:"comment:请求类型;"`
	RequestHeader  string `json:"requestHeader" gorm:"comment:请求头;"`
	ResponseHeader string `json:"responseHeader" gorm:"comment:响应头"`
	Status         bool   `json:"status" gorm:"comment:接口状态(0：关闭；1：开启)"`
}
