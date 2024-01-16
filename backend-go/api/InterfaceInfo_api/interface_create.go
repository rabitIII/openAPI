package InterfaceInfo_api

import (
	"backend-go/global"
	"backend-go/models"
	"backend-go/service/common/res"
	"github.com/gin-gonic/gin"
)

type InterfaceCreateRequest struct {
	Title          string `json:"title" binding:"required"` // 接口名称
	Description    string `json:"description"`              // 接口的描述
	Method         string `json:"method"`                   // 请求方式 get、post、...
	Url            string `json:"url"`                      // 接口的地址
	RequestHeader  string `json:"requestHeader"`            // 请求头
	ResponseHeader string `json:"responseHeader"`           // 响应头
}

func (InterfaceApi) InterfaceCreateView(c *gin.Context) {
	var cr InterfaceCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidError(err, &cr, c)
		return
	}

	var api models.InterfaceInfo
	global.DB.Where("title = ?", cr.Title).Or("url = ?", cr.Url).Find(&api)
	if err != nil {
		res.FailWithMsg("接口已存在", c)
		return
	}

	err = global.DB.Create(&models.InterfaceInfo{
		Title:          cr.Title,
		Description:    cr.Description,
		Method:         cr.Method,
		Url:            cr.Url,
		RequestHeader:  cr.RequestHeader,
		ResponseHeader: cr.ResponseHeader,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("接口创建失败", c)
		return
	}

	res.OKWithMsg("接口创建成功", c)

	return
}
