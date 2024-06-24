package core

import (
	"fmt"
	"log"
	"rabit-api-backend/global"
	"rabit-api-backend/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	dsn := global.Config.Databases.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[ERROR]mysql connect error: ", err.Error())
	}
	err = db.AutoMigrate(&models.UserModel{}, &models.InterfaceModel{})
	if err != nil {
		log.Fatal("[ERROR]create db tables error: ", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("[ERROR]获取sql实例失败！ error: ", err.Error())
		fmt.Println("[database mysql] 获取sql实例失败！")
		panic(err)
	}

	sqlDB.SetMaxIdleConns(global.Config.Databases.Mysql.MaxIdleConn)
	sqlDB.SetMaxOpenConns(global.Config.Databases.Mysql.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Hour * time.Duration(global.Config.Databases.Mysql.MaxLifeTime))

	fmt.Println("[SUCCESS] MySQL connect success! ")
	return db
}
