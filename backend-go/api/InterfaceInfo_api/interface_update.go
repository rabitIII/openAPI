package InterfaceInfo_api

import (
	"backend-go/global"
	"backend-go/models"
	"backend-go/service/common/res"
	"github.com/gin-gonic/gin"
)

type InterfaceUpdateRequest struct {
	ID             uint   `json:"ID"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Method         string `json:"method"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
}

func (InterfaceApi) InterfaceUpdateView(c *gin.Context) {
	var cr InterfaceUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		return
	}

	var api models.InterfaceInfo
	err = global.DB.Take(&api, "id = ?", cr.ID).Error
	if err != nil {
		res.FailWithMsg("接口不存在", c)
		return
	}

	err = global.DB.Model(&api).Updates(models.InterfaceInfo{
		Title:          cr.Title,
		Description:    cr.Description,
		Url:            cr.Url,
		Method:         cr.Method,
		RequestHeader:  cr.RequestHeader,
		ResponseHeader: cr.ResponseHeader,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("接口修改失败", c)
		return
	}

	res.OKWithMsg("接口修改成功", c)
	return
}
