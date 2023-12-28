package api

import "backend-go/api/user_api"

type Api struct {
	UserApi user_api.UserApi
}

var App = new(Api)
