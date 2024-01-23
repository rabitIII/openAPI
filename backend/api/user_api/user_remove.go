package user_api

import (
	"backend-go/global"
	"backend-go/models"
	"backend-go/service/common/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithError(err, c)
		return
	}

	var userList []models.UserModel
	global.DB.Find(&userList, cr.IDList)

	if len(cr.IDList) != len(userList) {
		res.FailWithMsg("数据一致性不通过", c)
		return
	}

	for _, model := range userList {
		global.DB.Delete(&model)
	}
	res.OKWithMsg(fmt.Sprintf("批量删除成功，共删除%d个用户", len(cr.IDList)), c)
	return
}
