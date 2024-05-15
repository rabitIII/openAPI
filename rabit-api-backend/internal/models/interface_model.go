package models

import (
	"gorm.io/gorm"
	"rabit-api-backend/global"
)

// InterfaceModel  -> database: `interface_info`
type InterfaceModel struct {
	gorm.Model
	Name           string            `gorm:"name"`
	Description    string            `gorm:"description"`
	Url            string            `gorm:"url"`
	Method         string            `gorm:"method"`
	RequestHeader  map[string]string `gorm:"requestHeader"`
	ResponseHeader map[string]string `gorm:"responseHeader"`
	UserId         int               `gorm:"userId"`
	Status         int               `gorm:"status"`
	IsDelete       int               `gorm:"isDelete"`
}

func (i *InterfaceModel) ReadList() error {
	// TODO 读取api接口列表
	//var interfaceModels []InterfaceModel
	// select * from interface_info
	conversation := "name,description,url,method,requestHeader,responseHeader"
	global.SqlDB.Raw("SELECT ? FROM interface_info", conversation)
	return nil
}
func (i *InterfaceModel) Create() error {
	global.SqlDB.Raw("INSERT INTO interface_info (name) VALUES (?)", i.Name)
	return nil
}

func (i *InterfaceModel) Update() error {
	return nil
}

func (i *InterfaceModel) Delete() error {
	return nil
}
