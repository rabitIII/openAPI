package router

import (
	"backend-go/global"
	"backend-go/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var apiGroup *gin.RouterGroup

// 初始化函数，即程序未执行前就开始初始化，若带有指针则会报错nil
//func init() {
//	r := gin.Default()
//
//	// 全局中间件
//	r.Use(middleware.Cors)
//	apiGroup = r.Group("api")
//
//	UserRouter()
//	//InterFaceRouter()
//}

func StartGin() {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.Cors, middleware.RefreshToken)
	apiGroup = r.Group("api")

	UserRouter()
	InterFaceRouter()
	addr := global.Config.System.Addr()
	err := r.Run(addr)
	if err != nil {
		fmt.Println("[Error] r.Run " + err.Error())
		return
	}
}
