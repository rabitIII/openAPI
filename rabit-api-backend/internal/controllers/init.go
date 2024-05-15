package controllers

import (
	"rabit-api-backend/internal/controllers/interface_controller"
	"rabit-api-backend/internal/controllers/user_controller"
)

// Controller 总控制器
type Controller struct {
	UserApi      user_controller.UserController
	InterfaceApi interface_controller.InterfaceController
}

var App = new(Controller)
