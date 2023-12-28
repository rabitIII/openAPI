package router

import (
	"backend-go/global"
	"backend-go/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var apiGroup *gin.RouterGroup

func init() {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.Cors)
	apiGroup = r.Group("api")

	addUserRouter()
}

func StartGin() {
	addr := global.Config.System.Addr()
	err := r.Run(addr)
	if err != nil {
		fmt.Println("[Error] r.Run " + err.Error())
		return
	}
}
