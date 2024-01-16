package router

import (
	"backend-go/api"
	"backend-go/middleware"
)

func UserRouter() {
	userApi := api.App.UserApi

	userRouter := apiGroup.Group("user")
	{
		/*
			login		登录
			register	注册（管理员进行创建）
			logout		登出
		*/
		userRouter.POST("login", userApi.UserLoginView)
		userRouter.POST("register", userApi.UserCreateView)
		//userRouter.POST("logout", app.UserLogoutView)

		/*
			currentUser	根据Token获取的id进行查找用户
		*/
		userRouter.GET("currentUser", middleware.RefreshToken, userApi.GetCurrentUser)
		userRouter.GET("list", userApi.UserListView)
	}
}
