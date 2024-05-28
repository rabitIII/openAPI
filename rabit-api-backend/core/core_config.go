package core

import (
	"fmt"
	"log"
	"os"
	"rabit-api-backend/config"

	"gopkg.in/yaml.v2"
)

const YAML_PATH = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(YAML_PATH)
	if err != nil {
		log.Fatal("[ERROR]read yaml error: ", err.Error())
	}

	c = new(config.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		log.Fatal("[ERROR]unmarshal yaml error: ", err.Error())
	}

	fmt.Println("[SUCCESS]read yaml OK!")
	return c
}
