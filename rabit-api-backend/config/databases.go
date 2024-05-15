package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rabit-api-backend/internal/models"
	"time"
)

var Conn *gorm.DB

type DatabaseConfig struct {
	Mysql Mysql `json:"mysql"`
}

type Mysql struct {
	Dsn         string `yaml:"dsn"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	MaxLifeTime int    `yaml:"maxLifeTime"`
}

func NewMySQL() {
	db, err := gorm.Open(mysql.Open(Conf.Databases.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.InterfaceModel{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(Conf.Databases.Mysql.MaxIdleConn)
	sqlDB.SetMaxOpenConns(Conf.Databases.Mysql.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Hour * time.Duration(Conf.Databases.Mysql.MaxLifeTime))

	Conn = db
}
