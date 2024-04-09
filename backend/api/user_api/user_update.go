package user_api

import (
	"backend-go/global"
	"backend-go/internal/models"
	"backend-go/internal/service/common/res"
	"backend-go/internal/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UserUpdateRequest struct {
	ID           uint   `json:"ID" binding:"required"`
	UserPassword string `json:"userPassword"`
	NickName     string `json:"nickName"`
	RoleID       uint   `json:"roleID"`
}

func (UserApi) UserUpdateView(c *gin.Context) {
	var cr UserUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		return
	}

	var api models.UserModel
	err = global.DB.Take(&api, "ID = ?", cr.ID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	if cr.UserPassword != "" {
		cr.UserPassword = pwd.HashPwd(cr.UserPassword)
	}
	err = global.DB.Model(&api).Updates(models.UserModel{
		UserPassword: cr.UserPassword,
		NickName:     cr.NickName,
		RoleID:       cr.RoleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户修改失败", c)
		return
	}

	res.OKWithMsg("用户修改成功", c)
	return
}
