package router

import "rabit-api-backend/internal/controllers"

func InterFaceRouter() {
	interfaceController := controllers.App.InterfaceApi

	interfaceRouter := indexGroup.Group("interface")
	{
		interfaceRouter.GET("/list")

	}
}
