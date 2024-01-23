package jwts

import (
	"backend-go/global"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenToken 生成token
func GenToken(user JwyPayLoad) (string, error) {
	//times := global.Config.Jwt.Expires
	cliams := CustomClaims{
		JwyPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                           // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                           // 生效时间
			//Issuer:    global.Config.Jwt.Issuer,                          // 签发人
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return t.SignedString([]byte(global.Config.Jwt.Secret))

}
