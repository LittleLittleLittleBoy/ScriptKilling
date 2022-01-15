package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/common/response"
	"github.com/littlelittlelittleboy/scriptkilling/common/vo"
	"github.com/littlelittlelittleboy/scriptkilling/constants"
)

type ScriptInitController struct{}

// create a new script.
func (controller *ScriptInitController) Create(c *gin.Context) {
	var request vo.CreateScriptRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		log.Fatal(err)
		response.GenerErrorResponse(c, nil, constants.REQUEST_PARAM_ERROR, "create script param error")
		return
	}

	if request.RoleNum <= 0 {
		response.GenerErrorResponse(c, nil, int(constants.SCRPIT_INFO_PARAM_ERROR), "role num less than 0")
		return
	}

	if len(request.Title) == 0 {
		response.GenerErrorResponse(c, nil, int(constants.SCRPIT_INFO_PARAM_ERROR), "script title can not be empty")
		return
	}

	scriptInfoService := ServiceInstance.ScriptInfoService
	count, result := scriptInfoService.Add(request)

	if result != constants.SCRIPT_INFO_SUCCESS {
		response.GenerErrorResponse(c, count, int(result), "insert script info error")
		return
	}

	response.GenerResponse(c, count)
}

// list all scripts.
func (controller *ScriptInitController) List(c *gin.Context) {

	scriptInfoService := ServiceInstance.ScriptInfoService
	data, result := scriptInfoService.List()
	if result != constants.SCRIPT_INFO_SUCCESS {
		response.GenerErrorResponse(c, nil, int(result), "list script info error")
		return
	}

	response.GenerResponse(c, data)
}

// add role in a script.
func (controller *ScriptInitController) AddPerson(c *gin.Context) {
	var request vo.CreateRoleRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		log.Fatal(err)
		response.GenerErrorResponse(c, nil, constants.REQUEST_PARAM_ERROR, "create role param error")
		return
	}

	if len(request.Name) == 0 {
		response.GenerErrorResponse(c, nil, int(constants.SCRPIT_ROLE_PARAM_ERROR), "create role name empty")
		return
	}

	scriptRoleService := ServiceInstance.ScriptRoleService
	count, result := scriptRoleService.Add(request)
	if result != constants.SCRIPT_ROLE_SUCCESS {
		response.GenerErrorResponse(c, count, int(result), "insert role info error")
		return
	}

	response.GenerResponse(c, count)
}

// list all persons
func (controller *ScriptInitController) ListAllPersons(c *gin.Context) {
	sidStr := c.Query("sid")
	if len(sidStr) == 0 {
		response.GenerErrorResponse(c, nil, constants.REQUEST_PARAM_ERROR, "query roles sid is empty")
		return
	}
	sid, error := strconv.Atoi(sidStr)
	if error != nil {
		response.GenerErrorResponse(c, nil, constants.REQUEST_PARAM_ERROR, "query roles sid is not int")
		return
	}
	scriptRoleService := ServiceInstance.ScriptRoleService
	data, result := scriptRoleService.List(sid)
	if result != constants.SCRIPT_ROLE_SUCCESS {
		response.GenerErrorResponse(c, data, int(result), "list script role error")
		return
	}
	response.GenerResponse(c, data)
}

// add a resource
func (controller *ScriptInitController) AddResource(c *gin.Context) {

}

// list
func (controller *ScriptInitController) ListAllResources(c *gin.Context) {

}
