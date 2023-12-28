package middleware

import "github.com/gin-gonic/gin"

func RefreshToken(c *gin.Context) {
	// 从请求头中获取Token，没有token就直接返回
	token := c.GetHeader("Authorization")
	if token == "" {
		c.Next()
	}

}
