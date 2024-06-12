package router

import (
	"rabit-api-backend/api"
)

func UserRouter() {
	userApi := api.RootApi.UserApi

	userRouter := indexGroup.Group("users")
	{
		/*
			login		用戶登录
			register	用戶注冊
			logout		退出登录

		*/
		userRouter.POST("login", userApi.UserLoginView)
		userRouter.POST("register", userApi.UserCreateView)
		userRouter.POST("logout", userApi.UserLogoutView)

		/*
			(管理员)
			list/:num/:page		获取全部的用户(分页)
			search/:id			获取单个用户
			delete/:id			删除指定用户
			update/:id			更新指定用户
			change/:id			修改指定用户的密码
		*/
		// userRouter.GET("list/:num/:page", middleware.AuthRole, userApi.UserListView)
		// userRouter.GET("search/:id", middleware.AuthRole, userApi.UserByIdView)
		// userRouter.POST("delete/:id", middleware.AuthRole, userApi.DeleteUserView)
		// userRouter.PUT("update/:id", middleware.AuthRole, userApi.UpdateUserView)
		// userRouter.PUT("change/:id", middleware.AuthRole, userApi.ChangePasswordView)
		// text
		userRouter.GET("list/:num/:page", userApi.UserListView)
		userRouter.GET("search/:id", userApi.UserByIdView)
		userRouter.DELETE("delete/:id", userApi.DeleteUserView)
		userRouter.PUT("update/:id", userApi.UpdateUserView)
		userRouter.PUT("change/:id", userApi.ChangePasswordView)
	}
}
