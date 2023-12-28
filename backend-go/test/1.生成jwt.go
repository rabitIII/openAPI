package main

import (
	"backend-go/utils/jwts"
	"fmt"
	"os"
)

func main() {
	var secretKey string = "qwertyuiop"
	token, err := jwts.GenToken(jwts.JwyPayLoad{
		NickName: "fengfeng",
		RoleID:   2,
		UserID:   1,
	}, secretKey)
	if err != nil {
		fmt.Println("生成jwt失败，", err)
		os.Exit(1)
	}
	fmt.Printf("%s", token)
}
