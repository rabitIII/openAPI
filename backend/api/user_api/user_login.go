package user_api

import (
	"backend-go/global"
	models2 "backend-go/internal/models"
	"backend-go/internal/service/common/res"
	jwts2 "backend-go/internal/utils/jwts"
	"backend-go/internal/utils/pwd"
	//"context"
	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"
)

// UserLoginRequest 登录时所需要的数据
type UserLoginRequest struct {
	UserAccount  string `json:"userAccount" binding:"required" label:"用户名"`
	UserPassword string `json:"userPassword" binding:"required" label:"密码"`
}

// UserLoginView
//
// Description: 用户登录，需要参数：userName、password
func (UserApi) UserLoginView(c *gin.Context) {
	var cr UserLoginRequest
	//var userDTO models.UserMode
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidError(err, &cr, c)
		return
	}

	// 数据库查询验证该账号数据
	var user models2.UserModel
	err = global.DB.Take(&user, "userAccount = ?", cr.UserAccount).Error
	if err != nil {
		global.Log.Warn("用户名不存在", cr.UserAccount)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !pwd.CheckPwd(user.UserPassword, cr.UserPassword) {
		global.Log.Warn("密码错误", cr.UserAccount, cr.UserPassword)
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
	token, err := jwts2.GenToken(jwts2.JwyPayLoad{
		NickName: user.UserAccount,
		RoleID:   user.RoleID,
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("生成token失败", c)
	}

	// 创建登录表
	err = global.DB.Create(&models2.LoginModel{
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

func (UserApi) UserLogoutView(c *gin.Context) {

}
