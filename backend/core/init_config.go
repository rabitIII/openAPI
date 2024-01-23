package core

import (
	"backend-go/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

const yamlpath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlpath)
	if err != nil {
		logrus.Fatalln("read yaml err: ", err.Error())
	}
	c = new(config.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		logrus.Fatalln("解析 yaml 失败: ", err.Error())
	}
	return c
}
