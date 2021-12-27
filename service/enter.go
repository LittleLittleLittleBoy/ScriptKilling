package service

type Service struct {
	UserService UserService
	FileService FileService
}

var ServiceInstance = new(Service)
