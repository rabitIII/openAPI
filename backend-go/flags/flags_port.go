package flags

import (
	"backend-go/global"
)

func Port(port int) {
	global.Config.System.Port = port
	//fmt.Println("修改程序运行的端口")
}
