package router

import "backend-go/api"

func addUserRouter() {
	app := api.App.UserApi

	userRouter := apiGroup.Group("user")
	{
		/*
			login		登录
			register	注册（管理员进行创建）
			logout		登出
		*/
		userRouter.POST("login", app.UserLoginView)
		userRouter.POST("register", app.UserCreateView)
		//userRouter.POST("logout", app.UserLogout)

		/*
			LoginToken
		*/
		userRouter.GET("loginToken", app.GetLoginUserUsingGet)
	}
}
