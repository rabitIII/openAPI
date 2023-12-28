package main

import (
	"backend-go/core"
	"backend-go/global"
	"backend-go/utils/jwts"
	"fmt"
	"os"
)

func main() {

	// 日志
	global.Log = core.InitLogger()
	// gin环境配置
	global.Config = core.InitConfig()
	s, err := jwts.GenToken(jwts.JwyPayLoad{
		NickName: "fengfeng",
	})
	if err != nil {
		fmt.Println("generate jwt failed, ", err)
		os.Exit(1)
	}
	fmt.Printf("token为：%s\n", s)

	// 解析jwt
	claims, err := jwts.ParseToken(s)
	if err != nil {
		fmt.Println("parse jwt failed, ", err)
		os.Exit(1)
	}
	fmt.Printf("token的解析为：%+v\n", claims)
}
