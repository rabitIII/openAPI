package router

import (
	"fmt"
	"net/http"
	"rabit-api-backend/global"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

var indexGroup *gin.RouterGroup

func init() {
	r = gin.Default()

	// main router
	indexGroup = r.Group("rabitApi")
	indexGroup.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
		})
	})

	// children router
	UserRouter()
}

func SetupGin() {
	addr := global.Config.Server.Http.Addr

	if err := r.Run(addr); err != nil {
		fmt.Println("[ERROR] running err: " + err.Error())
		return
	}
}
