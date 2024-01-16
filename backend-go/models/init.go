package models

import (
	"time"
)

//type LocalTime time.Time
//
//func (t *LocalTime) MarshalJSON() ([]byte, error) {
//	tTime := time.Time(*t)
//	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
//}

type Model struct {
	ID        uint      `json:"id" gorm:"primarykey"`            // 主键
	CretedAt  time.Time `gorm:"autoUpdateTime" json:"cretedAt"`  // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"` // 修改时间
}

type Pagination struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"Limit"`
	Key   string `json:"key" form:"key"`
	Sort  string `json:"sort" form:"sort"`
}

type IDListRequest struct {
	IDList []uint `json:"idList" form:"idList" binding:"required" label:"id列表"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}
