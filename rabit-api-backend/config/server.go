package config

type ServerConfig struct {
	Http Http `yaml:"http"`
}

type Http struct {
	Addr string `yaml:"addr"`
}
