package main

import (
	"backend-go/internal/conf"
	"backend-go/internal/db"
	"backend-go/internal/router"
)

func main() {

	conf.InitConfig()

	// 数据库初始化
	// db: MySQL, Redis
	db.InitDatabases()
	defer func() {
		db.Close()
	}()

	router.StartGin()
}
