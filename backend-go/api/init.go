package api

import (
	"backend-go/api/InterfaceInfo_api"
	"backend-go/api/user_api"
)

type Api struct {
	UserApi      user_api.UserApi
	InterfaceApi InterfaceInfo_api.InterfaceApi
}

var App = new(Api)
