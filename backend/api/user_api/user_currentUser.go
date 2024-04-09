package user_api

import (
	"backend-go/global"
	"backend-go/internal/models"
	"backend-go/internal/service/common/res"
	"github.com/gin-gonic/gin"
)

// GetUsrRequest 记录用户当前状态的属性
type GetUsrRequest struct {
	Token string `json:"token"`
	Data  any    `json:"data"`
}

func (UserApi) GetCurrentUser(c *gin.Context) {
	var cr GetUsrRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 检索用户
	var user models.LoginModel
	err = global.DB.Take(&user, "token = ?", cr.Token).Error
	if err == nil {
		res.FailWithMsg("未登录", c)
		return
	}

	res.OKWithMsg("用户创建成功", c)
}
