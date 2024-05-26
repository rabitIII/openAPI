package router

import "rabit-api-backend/api"

func UserRouter() {
	userApi := api.RootApi.UserApi

	userRouter := indexGroup.Group("user")
	{
		userRouter.POST("register", userApi.UserCreateView)
	}
}
