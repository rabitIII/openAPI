package router

import (
	"rabit-api-backend/api"
)

func InterfaceRouter() {
	interfaceApi := api.RootApi.InterfaceApi

	interfaceRouter := indexGroup.Group("interface")
	{
		interfaceRouter.POST("/create", interfaceApi.InterfaceCreateView)

	}
}
