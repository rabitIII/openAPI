package userapi

import (
	"fmt"

	"net/http"
	"rabit-api-backend/global"
	"rabit-api-backend/models"
	"rabit-api-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserLoginView
//
//	@Description: 用户登录，需要参数：user_account、user_password
//	@param c
func (UserApi) UserLoginView(ctx *gin.Context) {
	var userDTO models.UserInfo
	var user models.UserModel
	// 获取用户登录信息, 同时校验是否为空, 以及长度是否合法
	result := bindContextJson(ctx, &user)
	if !result {
		return
	}

	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched := utils.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, user.UserAccount)
	if !matched {
		ctx.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不合法！"))
		return
	}
	// 加密密码
	password := encryptString(user.UserPassword)

	// 查询数据库中是否存在该用户，并且同时把取出来的数据存入userDTO中
	affected := global.DB.Take(&models.UserModel{},
		"user_account = ? and user_password = ?", user.UserAccount, password).
		Scan(&userDTO).RowsAffected
	if affected == 0 {
		ctx.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不存在！"))
		return
	}

	// 记录用户的登录状态, 使用redis+token
	token := uuid.NewString()
	tokenKey := utils.TokenPrefix + token

	// 存入redis, 并且把用户ip存入redis
	err := global.Rdb.HSet(ctx, tokenKey, "id", userDTO.ID, "client_ip", ctx.ClientIP()).Err()

	if err != nil {
		fmt.Println("[api UserLogin err] Conn.Do HSET : " + err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "存储Token失败！"))
		return
	}
	// 设置有效期
	err = global.Rdb.Expire(ctx, tokenKey, utils.TokenTimeout).Err()
	if err != nil {
		fmt.Println("[api UserLogin err] Conn.Do EXPIRE : " + err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "设置Token有效期失败！"))
		return
	}

	// 封装user和token
	res := &gin.H{
		"user":  userDTO,
		"token": token,
	}
	ctx.JSON(http.StatusOK, utils.ResponseOK(res))
}

// UserLogoutView
//
//	@Description: 用户登出,无需前端传参数,将从token中取得当前的用户
//	@param ctx
func (UserApi) UserLogoutView(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	tokenKey := utils.TokenPrefix + token
	if token != "" {
		err := global.Rdb.HDel(ctx, tokenKey, "client_ip", "id").Err()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "用户Token已失效！"))
		}
	}
	ctx.JSON(http.StatusOK, utils.ResponseOK(nil))
}

// GetCurrentUser
//
// @Description 用户登录状态，
func (UserApi) GetCurrentUser(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "获取用户信息失败！"))
		return
	}

	var user models.UserInfo
	err := global.DB.Take(&models.UserModel{}, userId).Scan(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "用户不存在!"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(user))
}
