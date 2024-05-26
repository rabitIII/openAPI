package router

func InterfaceRouter() {

	interfaceRouter := indexGroup.Group("interface")
	{
		interfaceRouter.GET("/list")

	}
}
