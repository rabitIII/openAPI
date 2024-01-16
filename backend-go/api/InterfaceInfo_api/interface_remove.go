package InterfaceInfo_api

import (
	"backend-go/models"
	"github.com/gin-gonic/gin"
)

func (InterfaceApi) UserRemoveView(c *gin.Context) {
	var cr models.IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		return
	}

	//var apiList []models.InterfaceInfo
	//for _, apiId := range apiList
}
