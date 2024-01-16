package router

import "backend-go/api"

func InterFaceRouter() {
	interfaceApi := api.App.InterfaceApi

	interfaceRouter := apiGroup.Group("interface")
	{
		/*
			list	读取接口列表
			create	接口创建
			update	修改接口
		*/
		interfaceRouter.GET("/list", interfaceApi.InterfaceListView)
		interfaceRouter.POST("/create", interfaceApi.InterfaceCreateView)
		interfaceRouter.PUT("/update", interfaceApi.InterfaceUpdateView)
	}
}
