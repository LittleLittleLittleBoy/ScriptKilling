package controller

import (
	"log"

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
		response.GenerErrorResponse(c, nil, constants.SCRPIT_INFO_PARAM_ERROR, "role num less than 0")
		return
	}

	if len(request.Title) == 0 {
		response.GenerErrorResponse(c, nil, constants.SCRPIT_INFO_PARAM_ERROR, "script title can not be empty")
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

}

// list all persons
func (controller *ScriptInitController) ListAllPersons(c *gin.Context) {

}

// add a resource
func (controller *ScriptInitController) AddResource(c *gin.Context) {

}

// list
func (controller *ScriptInitController) ListAllResources(c *gin.Context) {

}
