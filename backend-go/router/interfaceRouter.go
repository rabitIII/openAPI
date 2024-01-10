package router

import "backend-go/api"

func InterFaceRouter() {
	interfaceApi := api.App.InterfaceApi

	interfaceRouter := apiGroup.Group("interface")
	{
		interfaceRouter.POST("/create", interfaceApi.InterfaceCreateView)
	}
}
