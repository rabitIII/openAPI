package utils

type BusCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Process StatusCode
var (
	SUCCESS = BusCode{0, "成功!"}
	WAIT    = BusCode{2, "等待处理中..."}
	EXPIRED = BusCode{7, "已过期!"}
)

// Error StatusCode
var (
	ParmasError = BusCode{50100, "请求参数错误!"}
	LoginError  = BusCode{20100, "未登录!"}
	AuthError   = BusCode{40100, "未授权!"}
	ServerError = BusCode{50000, "系统错误!"}
	MysqlError  = BusCode{50001, "Mysql错误!"}
	RedisError  = BusCode{50002, "Redis错误!"}
)

// User StatusCode
var (
	RegisterOK = BusCode{20000, "用户注册成功!"}
	LoginOK    = BusCode{20100, "用户登录成功!"}
	LogoutOK   = BusCode{20200, "用户登出!"}
	UpdateOK   = BusCode{20300, "用户信息更新成功!"}
	DeleteOK   = BusCode{20100, "用户账号已注销!"}
)
