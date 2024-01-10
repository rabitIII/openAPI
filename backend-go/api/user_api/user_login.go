package user_api

import (
	"backend-go/global"
	"backend-go/models"
	"backend-go/service/common/res"
	"backend-go/utils/jwts"
	"backend-go/utils/pwd"
	//"context"
	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"
)

// UserLoginRequest 登录时所需要的数据
type UserLoginRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

// UserLoginView
//
// Description: 用户登录，需要参数：userName、password
func (UserApi) UserLoginView(c *gin.Context) {
	var cr UserLoginRequest
	//var userDTO models.UserModel
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 数据库查询验证该账号数据
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在", cr.UserName)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !pwd.CheckPwd(user.Password, cr.Password) {
		global.Log.Warn("密码错误", cr.UserName, cr.Password)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	//// 记录用户的登录状态，使用redis+token
	//token := uuid.NewString()
	//fmt.Println(token)
	//tokenKey := jwts.TokenPrefix + token
	//// 存入redis，并且把用户ip存入redis中
	////var ctx context.Context = context.Background()
	//err = global.Redis.HSet(tokenKey, c.ClientIP(), "id").Err()
	//
	//if err != nil {
	//	logrus.Fatal("[api UserLogin err] Conn.Do HSET : ", err.Error())
	//	c.JSON(500, "存储失败")
	//	return
	//}
	//fmt.Println(c.ClientIP())
	////设置token有效期
	//err = global.Redis.Expire(tokenKey, time.Duration(global.Config.Jwt.Expires)).Err()
	//if err != nil {
	//	logrus.Fatal("[api UserLogin err] Conn.Do EXPIRE : ", err.Error())
	//	return
	//}
	//
	//res := &gin.H{
	//	"token": token,
	//}
	//c.JSON(200, utils.ResponseOK(res))

	// 验证通过后将response里的数据进行加密
	token, err := jwts.GenToken(jwts.JwyPayLoad{
		NickName: user.UserName,
		RoleID:   user.RoleID,
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("生成token失败", c)
	}

	// 创建登录表
	err = global.DB.Create(&models.LoginModel{
		UserID:   user.Model.ID,
		NickName: user.NickName,
		Token:    token,
		RoleID:   user.RoleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		//logrus.Fatal(err)
		res.FailWithMsg("用户登录失败", c)
		return
	}

	// response返回
	res.OKWithData(token, c)
	return
}
