package config

import "fmt"

type ServerConfig struct {
	Http Http `yaml:"http"`
}

type Http struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

func (h *Http) Addr() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}
