package user_api

import (
	"backend-go/global"
	"backend-go/internal/models"
	"backend-go/internal/service/common/res"
	"backend-go/internal/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserAccount  string `json:"userAccount" binding:"required"`  // 用户名
	UserPassword string `json:"userPassword" binding:"required"` // 密码
	NickName     string `json:"nickName"`                        // 昵称
	RoleID       uint   `json:"roleID" binding:"required"`       // 角色id,权限表的id
	//IsSystem bool   `json:"isSystem" binding:"required"` // 系统权限(4:普通用户,1:管理员)
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidError(err, &cr, c)
		return
	}

	// 检索用户
	var user models.UserModel
	err = global.DB.Take(&user, "userAccount = ?", cr.UserAccount).Error
	if err == nil {
		res.FailWithMsg("账号存在", c)
		return
	}

	if cr.NickName == "" {
		// 用户的昵称不存在，系统创建 user_ + "id"
		var maxId uint
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxId)
		cr.NickName = fmt.Sprintf("user_%d", maxId+1)
	}

	err = global.DB.Create(&models.UserModel{
		UserAccount:  cr.UserAccount,
		UserPassword: pwd.HashPwd(cr.UserPassword),
		NickName:     cr.NickName,
		RoleID:       cr.RoleID,
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
