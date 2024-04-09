package jwts

import "github.com/golang-jwt/jwt/v5"

type JwyPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
}

var Secret []byte

type CustomClaims struct {
	JwyPayLoad
	jwt.RegisteredClaims
}
