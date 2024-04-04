package server

import (
	"backend-go/config"
	"backend-go/db"
)

func Start() {
	config.InitConfig()
	defer func() {
		db.Close()
	}()
}
