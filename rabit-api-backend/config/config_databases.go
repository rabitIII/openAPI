package config

import "fmt"

type DatabaseConfig struct {
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	DBname      string `yaml:"dbname"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	MaxLifeTime int    `yaml:"maxLifeTime"`
}

func (m *Mysql) Dsn() string {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DBname)
}
