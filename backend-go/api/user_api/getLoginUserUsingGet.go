package user_api

import (
	"backend-go/service/common/res"
	"github.com/gin-gonic/gin"
)

//// GetUsrRequest 记录用户当前状态的属性
//type GetUsrRequest struct {
//	//Token string `json:"token"`
//	Data any `json:"data"`
//}

func (UserApi) GetLoginUserUsingGet(c *gin.Context) {
	var cr UserLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	res.OKWithData(data, c)
	return
}
