package user_api

import (
	"backend-go/global"
	"backend-go/models"
	"backend-go/service/common/res"
	"backend-go/utils/jwts"
	"backend-go/utils/pwd"
	"github.com/gin-gonic/gin"
)

// UserLoginRequest 登录时所需要的数据
type UserLoginRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

func (UserApi) UserLoginView(c *gin.Context) {
	var cr UserLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 数据库查询验证该账号数据
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在", cr.UserName)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !pwd.CheckPwd(user.Password, cr.Password) {
		global.Log.Warn("密码错误", cr.UserName, cr.Password)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	// 验证通过后将response里的数据进行加密
	token, err := jwts.GenToken(jwts.JwyPayLoad{
		NickName: user.UserName,
		RoleID:   user.RoleID,
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("生成token失败", c)
	}

	// response返回
	res.OKWithData(token, c)
	return
}
