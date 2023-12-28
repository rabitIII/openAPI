package user_api

import (
	"backend-go/global"
	"backend-go/models"
	"backend-go/service/common/res"
	"backend-go/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	NickName string `json:"nickName"`                    // 昵称
	RoleID   uint   `json:"roleID" binding:"required"`   // 角色id
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 检索用户
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err == nil {
		res.FailWithMsg("用户名已存在", c)
		return
	}

	if cr.NickName == "" {
		// 用户的昵称不存在，系统创建 user_ + "id"
		var maxId uint
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxId)
		cr.NickName = fmt.Sprintf("user_%d", maxId+1)
	}

	err = global.DB.Create(&models.UserModel{
		UserName: cr.UserName,
		Password: pwd.HashPwd(cr.Password),
		NickName: cr.NickName,
		RoleID:   cr.RoleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		//logrus.Fatal(err)
		res.FailWithMsg("用户创建失败", c)
		return
	}

	res.OKWithMsg("用户创建成功", c)

	return
}
