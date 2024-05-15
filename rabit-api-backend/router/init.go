package router

import "github.com/gin-gonic/gin"

var indexGroup *gin.RouterGroup

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	indexGroup = r.Group("rabit-api")
	{
		InterFaceRouter()
	}
	return r
}
