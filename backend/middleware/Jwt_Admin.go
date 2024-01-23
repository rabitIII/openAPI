package middleware

import (
	"backend-go/service/common/res"
	"backend-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

func JwtAdmin(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		res.FailWithMsg("未携带token", c)
		c.Abort()
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMsg("token错误", c)
		c.Abort()
		return
	}

	if claims.RoleID != 1 {
		res.FailWithMsg("权限错误！", c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
}
