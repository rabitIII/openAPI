package models

import (
	"gorm.io/gorm"
)

// InterfaceModel  -> database: `interface_info`
type InterfaceModel struct {
	gorm.Model
	Name           string `gorm:"name"`
	Description    string `gorm:"description"`
	Url            string `gorm:"url"`
	Method         string `gorm:"method"`
	RequestHeader  string `gorm:"request_header"`
	ResponseHeader string `gorm:"response_header"`
	UserId         uint   `gorm:"user_id"`
	Status         bool   `gorm:"status"`
	IsDelete       bool   `gorm:"isDelete"`
}

type InterfaceCreateRequest struct {
	UserId         uint   `json:"user_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	Url            string `json:"url" binding:"required"`
	Method         string `json:"method" binding:"required"`
	RequestHeader  string `json:"request_header"`
	ResponseHeader string `json:"response_header"`
	RequestParams  string `json:"request_params"`
}
