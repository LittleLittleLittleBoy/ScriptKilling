package route

import (
	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/common/response"
	"github.com/littlelittlelittleboy/scriptkilling/controller"
)

type ScriptCreatorRouter struct{}

func (r *ScriptCreatorRouter) Init(group *gin.RouterGroup) {
	apiGroup := group.Group("/api/admin")
	{
		fileUpload := controller.Instance.FileUpload
		apiGroup.POST("upload", fileUpload.Upload)
	}
	apiGroup.GET("create", func(c *gin.Context) {
		response.GenerResponse(c, "ok")
	})
}
