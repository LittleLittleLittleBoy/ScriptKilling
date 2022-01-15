package service

type Service struct {
	UserService       UserService
	FileService       FileService
	ScriptInfoService ScriptInfoService
	ScriptRoleService ScriptRoleService
}

var ServiceInstance = new(Service)
