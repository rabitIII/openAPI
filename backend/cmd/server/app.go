package main

import (
	"backend-go/config"
	"backend-go/db"
)

func main() {

	config.InitConfig()

	// 数据库初始化
	// db: MySQL, Redis
	db.InitDatabases()
	defer func() {
		db.Close()
	}()
}
