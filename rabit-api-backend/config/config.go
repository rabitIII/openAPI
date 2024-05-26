package config

type Config struct {
	Server    ServerConfig   `yaml:"server"`
	Databases DatabaseConfig `yaml:"databases"`
}
