package user_api

import (
	"backend-go/models"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithError(err, c)

	}
}
