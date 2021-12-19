package controller

import "github.com/littlelittlelittleboy/scriptkilling/service"

type Controller struct {
	User UserController
}

var Instance = new(Controller)

var userService = service.ServiceInstance.UserService
