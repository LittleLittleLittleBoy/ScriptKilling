package service

type Service struct {
	UserService       UserService
	FileService       FileService
	ScriptInfoService ScriptInfoService
}

var ServiceInstance = new(Service)
