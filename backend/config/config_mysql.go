package config

type Mysql struct {
	Host     string `yaml:"host"`     // 服务器地址url
	Port     string `yaml:"port"`     // 端口
	Config   string `yaml:"config"`   // 高级配置
	DB       string `yaml:"db"`       // 数据库名
	Username string `yaml:"username"` // 数据库登录名
	Password string `yaml:"password"` // 数据库登录密码
	LogLevel string `yaml:"LogLevel"` // 是否开启Gorm全局日志
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Port + ")/" + m.DB + "?" + m.Config
}
