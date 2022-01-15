package service

type Service struct {
	UserService              UserService
	FileService              FileService
	ScriptInfoService        ScriptInfoService
	ScriptRoleService        ScriptRoleService
	ScriptInformationService ScriptInformationService
}

var ServiceInstance = new(Service)
