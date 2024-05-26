package main

import (
	"rabit-api-backend/core"
	"rabit-api-backend/global"
	"rabit-api-backend/router"
)

func main() {
	// 初始化配置
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()

	router.SetupGin()
}
