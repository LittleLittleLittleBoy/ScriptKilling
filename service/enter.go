package service

type Service struct {
	UserService UserService
}

var ServiceInstance = new(Service)
