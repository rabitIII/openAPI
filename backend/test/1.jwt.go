package test

import (
	"backend-go/utils/jwts"
	"fmt"
	"os"
	"testing"
)

// TestJwt 创建jwt
func TestJwt(t *testing.T) {
	//var secretKey string = "qwertyuiop"
	//token, err := jwts.GenToken(jwts.JwyPayLoad{
	//	NickName: "fengfeng",
	//	RoleID:   2,
	//	UserID:   1,
	//}, secretKey)
	token, err := jwts.GenToken(jwts.JwyPayLoad{
		NickName: "fengfeng",
		RoleID:   2,
		UserID:   1,
	})
	if err != nil {
		fmt.Println("生成jwt失败，", err)
		os.Exit(1)
	}
	fmt.Printf("%s", token)
}
