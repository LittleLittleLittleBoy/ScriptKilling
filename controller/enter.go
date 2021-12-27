package controller

import "github.com/littlelittlelittleboy/scriptkilling/service"

type Controller struct {
	User       UserController
	ScriptInit ScriptInitController
	FileUpload FileUploadControler
}

var Instance = new(Controller)

var ServiceInstance = service.ServiceInstance
