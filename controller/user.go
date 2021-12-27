package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/common/response"
	"github.com/littlelittlelittleboy/scriptkilling/constants"
)

type UserController struct{}

func (controller *UserController) Login(c *gin.Context) {
	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))

	if username == "" || password == "" {
		response.GenerErrorResponse(c, nil, int(constants.LOGIN_EMPTY_INFO), "Empty username or password")
		return
	}

	userService := ServiceInstance.UserService

	status := userService.Login(username, password)

	if status != constants.LOGIN_SUCCESS {
		response.GenerErrorResponse(c, nil, int(status), "Login error")
		return
	}

	response.GenerResponse(c, "Success")
}

func (controller *UserController) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		response.GenerErrorResponse(c, nil, int(constants.LOGIN_EMPTY_INFO), "Empty username or password")
		return
	}

	userService := ServiceInstance.UserService

	status := userService.Register(username, password)

	if status != constants.LOGIN_SUCCESS {
		response.GenerErrorResponse(c, nil, int(status), "Register error")
		return
	}

	response.GenerResponse(c, "Success")
}
