package userapi

import (
	"fmt"
	"log"
	"net/http"
	"rabit-api-backend/global"
	"rabit-api-backend/models"

	"github.com/gin-gonic/gin"
)

type Responce struct {
	Code int
	Data int
	Msg  string
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr models.UserCreateRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		log.Fatalln("[ERROR] 读参错误, err:%w", err)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "userAccount = ?", cr.UserAccount).Error
	if err == nil {
		log.Fatalln("账号存在")
		return
	}

	if cr.NickName == "" {
		var maxId uint

		// SQL: select max(id) from user order by
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxId)
		cr.NickName = fmt.Sprintf("user_%d", maxId+1)
	}

	err = global.DB.Create(&models.UserModel{
		UserAccount:  cr.UserAccount,
		UserPassword: cr.UserPassword,
		NickName:     cr.NickName,
		RoleID:       cr.RoleID,
	}).Error
	if err != nil {
		log.Fatalln("用户创建失败, err:%w", err)
		return
	}

	c.JSON(http.StatusOK, Responce{0, 1, "created success!"})
}
