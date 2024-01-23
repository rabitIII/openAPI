package main

import (
	"backend-go/core"
	"backend-go/flags"
	"backend-go/global"
	"backend-go/router"
)

func main() {
	// 日志
	global.Log = core.InitLogger()
	// gin环境配置
	global.Config = core.InitConfig()
	// mysql初始化
	global.DB = core.InitMysql()
	// redis
	global.Redis = core.InitRedis(0)

	//val, err := global.Redis.Get("name").Result()
	//fmt.Println(val, err)

	// 命令行参数运行脚本
	option := flags.Parse()
	if option.Run() {
		return
	}

	//r := router.Router()
	////fmt.Println(global.Config)
	//addr := global.Config.System.Addr()
	//r.Run(addr)
	router.StartGin()
}
