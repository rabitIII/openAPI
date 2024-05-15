package interface_controller

import "github.com/gin-gonic/gin"

func (InterfaceController) GetInterface(c *gin.Context) {
	c.JSON(200, gin.H{})
}
