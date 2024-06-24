package api

import (
	interfaceapi "rabit-api-backend/api/interfaceApi"
	userapi "rabit-api-backend/api/userApi"
)

type Api struct {
	UserApi      userapi.UserApi
	InterfaceApi interfaceapi.InterfaceApi
}

var RootApi = new(Api)
