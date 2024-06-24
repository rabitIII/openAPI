package interfaceapi

import (
	"fmt"
	"net/http"
	"rabit-api-backend/global"
	"rabit-api-backend/models"
	"rabit-api-backend/utils"

	"github.com/gin-gonic/gin"
)

func (InterfaceApi) InterfaceCreateView(c *gin.Context) {

	var cr models.InterfaceCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		fmt.Println("[ERROR] CreateInterfaceInfo: requestParams error!", err.Error())
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "添加新的接口的参数有误！"))
		return
	}

	var interfaceInfo = &models.InterfaceModel{
		Name:           cr.Name,
		UserId:         cr.UserId,
		Description:    cr.Description,
		Url:            cr.Url,
		Method:         cr.Method,
		RequestHeader:  cr.RequestHeader,
		ResponseHeader: cr.ResponseHeader,
		Status:         false,
	}

	affected := global.DB.Save(interfaceInfo).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "创建接口失败！"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseOK(cr))
}
