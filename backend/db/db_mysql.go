package db

import (
	"backend-go/config"
	"backend-go/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Conn *gorm.DB

func InitMysql() {
	if config.Conf.Mysql.Host == "" {
		panic("[ERROR] 未配置mysql，取消gorm连接")
	}
	dsn := global.Config.Mysql.Dsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error(fmt.Sprintf("[%s] mysql 连接失败，err:%s", dsn, err.Error()))
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout

	Conn = db
}

func Close() {
	db, _ := Conn.DB()
	err := db.Close()
	if err != nil {
		return
	}
}
