package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Jwt    Jwt    `yaml:"jwt"`
}

var Conf Config

func InitConfig() {
	yamlPath, err := os.ReadFile("settings.yaml")
	if err != nil {
		logrus.Fatalln("[ERROR] read yaml err: ", err.Error())
	}
	err = yaml.Unmarshal(yamlPath, &Conf)
	if err != nil {
		logrus.Fatalln("[ERROR] 解析 yaml 失败: ", err.Error())
	}
	fmt.Println("[SUCCESS] ConfigFile read successfully！")
}
