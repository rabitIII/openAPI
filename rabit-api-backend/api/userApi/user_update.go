package userapi

import (
	"net/http"
	"rabit-api-backend/global"
	"rabit-api-backend/models"
	"rabit-api-backend/utils"

	"github.com/gin-gonic/gin"
)

// UpdateUserView
//
//	 @Description: 管理员修改指定id的用户, 需要请求参数: id
//					 并且需要JSON参数: id(必须),user_account(必须),user_name,gender,phone,email,avatar_url
//	 @param c
func (UserApi) UpdateUserView(c *gin.Context) {
	//验证id是否合法
	id := utils.IsNumber(c.Param("id"))
	if id == -1 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id错误！"))
		return
	}
	// 获取前端发送的用户信息
	var user models.UserInfo
	result := bindContextJson(c, &user)
	if !result {
		return
	}

	// 判断请求参数的id是否和获取到的用户id一致，不一致则直接返回
	if id != int(user.ID) {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "需要修改的用户参数和请求不合法！"))
		return
	}

	var userAccount string
	// 必须先获取用户原本的用户名，再进行判断用户是否已经更改了用户名
	affected := global.DB.Take(&models.UserModel{}, id).Select("user_account").Scan(&userAccount).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusForbidden,
			utils.ResponseError(utils.ParamsError, "请求的id用户不存在！"))
		return
	}
	// 用户名发生了变化
	if userAccount != user.UserAccount {
		// 判断更改后的用户名是否存在
		exist := isUserAccountExist(c, user.UserAccount)
		if exist {
			return
		}
	}

	// 更新用户的信息
	isSuccess := updateUserInfo(c, user)
	if !isSuccess {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "Mysql修改用户信息错误！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil, "修改用户信息成功！"))
}

// ChangePasswordView
//
//	 @Description: 管理员更改用户的密码,需要请求参数: id
//									需要JSON参数: id,user_password,check_password (三个参数均必须)
//	 @param c
func (UserApi) ChangePasswordView(c *gin.Context) {
	// 获取请求的id
	id := utils.IsNumber(c.Param("id"))
	if id == -1 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id错误！"))
		return
	}
	// 获取需要修改的用户id、用户名、和修改后的密码
	var user models.UserChangePassword
	isSuccess := bindContextJson(c, &user)
	if !isSuccess {
		return
	}

	// 判断请求的id和发送的用户数据中的id是否一致
	if id != int(user.ID) {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "请求的id和发送的id不一致！"))
		return
	}

	isSuccess = changePassword(c, user)
	if !isSuccess {
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil))
}
