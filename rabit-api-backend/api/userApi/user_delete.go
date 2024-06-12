package userapi

import (
	"net/http"
	"rabit-api-backend/global"
	"rabit-api-backend/models"
	"rabit-api-backend/utils"

	"github.com/gin-gonic/gin"
)

// DeleteUserById
//
//	@Description: 删除指定用户，需要接受请求参数: id
//	@param c
func (UserApi) DeleteUserView(c *gin.Context) {
	//验证id是否合法
	id := utils.IsNumber(c.Param("id"))
	if id == -1 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id错误！"))
		return
	}
	// 获取发起请求的用户id
	userId := getContextValue(c, "userId")
	if userId == nil {
		return
	}
	// 判断是否删除的用户是自己 如果删除的id和发起请求人的id一致则不能删除
	if userId == id {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "不能删除自己！"))
		return
	}
	// 删除用户
	affected := global.DB.Delete(&models.UserModel{}, id).RowsAffected
	// 用户已经不存在了
	if affected == 0 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "该用户已经不存在！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil))
}
