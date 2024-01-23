package InterfaceInfo_api

import (
	"backend-go/models"
	"backend-go/service/common/list"
	"backend-go/service/common/res"
	"github.com/gin-gonic/gin"
)

type InterfaceRequest struct {
	Page  int    `json:"page" form:"page"` // 页
	Limit int    `json:"limit" form:"limit"`
	Key   string `json:"key" form:"key"`
}

func (InterfaceApi) InterfaceListView(c *gin.Context) {

	var cr models.Pagination
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		return
	}

	_list, count, _ := list.QueryList(models.InterfaceInfo{}, list.Option{
		Pagination: cr,
		Likes:      []string{"title"},
		Debug:      true,
	})
	//var interfacelist []models.InterfaceInfo
	//
	//var count int                      // 总数
	//offset := (cr.Page - 1) * cr.Limit // 分页
	//
	//if cr.Limit <= 0 {
	//	cr.Limit = 10
	//}
	//
	//global.DB.Model(models.InterfaceInfo{}).Select("count(id)").Scan(&count)
	//global.DB.Limit(cr.Limit).Offset(offset).Find(&interfacelist)
	res.OKWithList(_list, count, c)
}
