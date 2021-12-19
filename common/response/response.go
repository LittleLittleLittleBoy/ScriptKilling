package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func GenerResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"",
	})
}

func GenerErrorResponse(c *gin.Context, data interface{}, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
