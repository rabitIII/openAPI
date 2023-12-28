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
