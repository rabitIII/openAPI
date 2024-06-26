package userapi

import (
	"fmt"
	"net/http"
	"rabit-api-backend/global"
	"rabit-api-backend/models"
	"rabit-api-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserListView 获取所有用户的信息列表
func (UserApi) UserListView(ctx *gin.Context) {
	// 获取num和page的参数，并验证
	numParam := ctx.Param("num")
	matched := utils.MatchString(`^[0-9]*$`, numParam)
	if !matched {
		ctx.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "num必须为数字！"))
		return
	}
	pageParam := ctx.Param("page")
	matched = utils.MatchString(`^[0-9]*$`, pageParam)
	if !matched {
		ctx.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "page必须为数字！"))
		return
	}

	// 将num和page转化为整数
	num, err := strconv.Atoi(numParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "转化为数字失败！"))
		return
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "转化为数字失败！"))
		return
	}

	var userList []models.UserInfo
	offset := (page - 1) * num
	affected := global.DB.Limit(num).Offset(offset).Model(&models.UserModel{}).Scan(&userList).RowsAffected
	if affected == 0 {
		ctx.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "数据库中没有用户！"))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseOK(userList, fmt.Sprintf("查找到了%v个用户！", affected)))
}

// UserByIdView	获取指定id的用户信息
func (UserApi) UserByIdView(ctx *gin.Context) {
	// 验证id是否合法
	id := utils.IsNumber(ctx.Param("id"))
	if id == -1 {
		ctx.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id错误！"))
		return
	}

	var userDTO models.UserInfo
	affected := global.DB.Take(&models.UserModel{}, id).Scan(&userDTO).RowsAffected
	if affected == 0 {
		ctx.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "该用户不存在！"))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResponseOK(userDTO))
}
