package api

import userapi "rabit-api-backend/api/userApi"

type Api struct {
	UserApi *userapi.UserApi
}

var RootApi = new(Api)
