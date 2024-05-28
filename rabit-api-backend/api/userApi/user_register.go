package userapi

import (
	"fmt"
	"net/http"
	"rabit-api-backend/global"
	"rabit-api-backend/models"
	"rabit-api-backend/services/res"

	"github.com/gin-gonic/gin"
)

func (UserApi) UserCreateView(ctx *gin.Context) {
	var cr models.UserCreateRequest

	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		fmt.Println("空指针: ", err.Error())
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "userAccount = ?", cr.UserAccount).Error
	if err == nil {
		ctx.JSON(http.StatusForbidden, res.ResponseError(res.ParamsError, "账号已存在！"))
		return
	}

	if cr.NickName == "" {
		var maxId uint

		// SQL: select max(id) from user order by
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxId)
		cr.NickName = fmt.Sprintf("user_%d", maxId+1)
	}

	var userDTO models.UserInfo

	userData := &models.UserModel{
		UserAccount:  cr.UserAccount,
		UserPassword: cr.UserPassword,
		NickName:     cr.NickName,
		RoleID:       cr.RoleID,
	}

	affected := global.DB.Create(userData).Scan(&userDTO).RowsAffected
	if affected == 0 {
		ctx.JSON(http.StatusInternalServerError, res.ResponseError(res.ServerError, "user register err"))
	}

	ctx.JSON(http.StatusOK, res.ResponseOK(userDTO, "create user success!"))
}
