package middleware

import (
	"fmt"
	"rabit-api-backend/global"
	"rabit-api-backend/utils"

	"github.com/gin-gonic/gin"
)

// RefreshToken
//
//	@Description: 全局中间件，判断有token就刷新，无token直接放行
//	@param c
func RefreshToken(c *gin.Context) {
	// 从请求头中获取Token, 没有token就直接返回
	token := c.GetHeader("Authorization")
	if token == "" {
		c.Next()
	}
	tokenKey := utils.TokenPrefix + token
	// 在redis中查找是否Token是否存在，允许未查到
	id, err := global.Rdb.HGet(c, tokenKey, "id").Result()
	if err != nil {
		fmt.Println("[middleware RefreshToken err] database.Client.HGet ", err.Error())
		return
	}
	// 没有查询到任何用户，直接放行
	if id == "" {
		c.Next()
	}
	// 查找到了用户，刷新token时间
	err = global.Rdb.Expire(c, tokenKey, utils.TokenTimeout).Err()
	if err != nil {
		fmt.Println("[middleware RefreshToken err] database.Client.Expire ", err.Error())
		return
	}
}
