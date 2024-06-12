package userapi

import (
	"net/http"
	"rabit-api-backend/global"
	"rabit-api-backend/models"
	"rabit-api-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserCreateView
//
// @Description 用户注册，创建新的用户
// @param ctx
func (UserApi) UserCreateView(ctx *gin.Context) {
	var cr models.UserCreateRequest

	isNull := bindContextJson(ctx, &cr)
	// err := ctx.ShouldBindJSON(&cr)
	if !isNull {
		return
	}

	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched := utils.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, cr.UserAccount)
	if !matched {
		ctx.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不合法！"))
		return
	}
	// 密码(以字母开头，只能包含字母、数字和下划线)：^[a-zA-Z]\w*$    \w = [a-zA-Z0-9_]

	// 账户是否重复
	exist := isUserAccountExist(ctx, cr.UserAccount)
	if exist {
		return
	}

	// 密码加密
	password := encryptString(cr.UserPassword)
	// accessKey生成
	accessKey := encryptString(cr.UserAccount + uuid.NewString())
	secretKey := encryptString(password + uuid.NewString())

	var userDTO models.UserInfo

	// 插入数据
	user := &models.UserModel{
		UserPassword: password,
		UserAccount:  cr.UserAccount,
		NickName:     cr.NickName,
		RoleID:       cr.RoleID,
		AccessKey:    accessKey,
		SecretKey:    secretKey,
	}
	affected := global.DB.Save(user).Scan(&userDTO).RowsAffected
	if affected == 0 {
		ctx.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "注册用户失败！"))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseOK(userDTO))
}
