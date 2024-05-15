package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server    ServerConfig   `yaml:"server"`
	Databases DatabaseConfig `yaml:"databases"`
}

var Conf Config

func InitConfig() {
	file, err := os.ReadFile("./../configs/settings.yaml")
	if err != nil {
		fmt.Println("[ERROR]read yaml error:", err)
		return
	}

	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		fmt.Println("[ERROR]unmarshal yaml error:", err)
	}
	fmt.Println(Conf)

}
