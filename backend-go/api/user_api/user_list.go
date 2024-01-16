package user_api

import (
	"backend-go/models"
	"backend-go/service/common/list"
	"backend-go/service/common/res"
	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Key   string `json:"key" form:"key"`
}

func (UserApi) UserListView(c *gin.Context) {

	var cr models.Pagination
	err := c.ShouldBindQuery(cr)
	if err != nil {
		return
	}

	_list, count, _ := list.QueryList(models.UserModel{}, list.Option{
		Pagination: cr,
		Likes:      []string{"nickName", "userName"},
		Preload:    []string{"RoleModel"},
		Debug:      true,
	})

	res.OKWithList(_list, count, c)

}
