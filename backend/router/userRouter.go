package router

import (
	"backend-go/api"
	"backend-go/middleware"
)

func UserRouter() {
	userApi := api.App.UserApi

	userRouter := apiGroup.Group("users")
	{
		/*
			login		登录
			register	注册（管理员进行创建）
			logout		登出
		*/
		userRouter.POST("login", userApi.UserLoginView)
		userRouter.POST("register", userApi.UserCreateView)
		//userRouter.POST("logout", userApi.UserLogoutView)

		/*
			currentUser	根据Token获取的id进行查找用户
		*/
		userRouter.GET("currentUser", middleware.RefreshToken, userApi.GetCurrentUser)

		/*
			list	用户列表
			update	修改用户信息
			delete	删除用户
		*/
		userRouter.GET("list", userApi.UserListView)
		userRouter.PUT("update", userApi.UserUpdateView)
		userRouter.POST("delete", userApi.UserRemoveView)
	}
}
