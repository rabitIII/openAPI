package flags

import (
	"backend-go/global"
	models2 "backend-go/internal/models"
	"github.com/sirupsen/logrus"
)

func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models2.UserModel{},
		&models2.RoleModel{},
		&models2.LoginModel{},
		&models2.LoginModel{},
		&models2.InterfaceInfo{},
	)

	if err != nil {
		logrus.Fatalf("生成数据库迁移失败，err:%s", err.Error())
	}
	logrus.Infof("数据库迁移成功！")
	//fmt.Println("修改数据库")
}
