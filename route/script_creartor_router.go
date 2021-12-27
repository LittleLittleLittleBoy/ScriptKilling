package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/constants"
	"github.com/littlelittlelittleboy/scriptkilling/controller"
)

type ScriptCreatorRouter struct{}

func (r *ScriptCreatorRouter) Init(group *gin.RouterGroup) {
	group.StaticFS(constants.FILE_BASE_PATH, http.Dir(constants.FILE_BASE_PATH))

	apiGroup := group.Group("/api/admin")
	{
		fileUpload := controller.Instance.FileUpload
		apiGroup.POST("upload", fileUpload.Upload)
	}
	{
		scriptInfo := controller.Instance.ScriptInit
		apiGroup.POST("scriptInfo/add", scriptInfo.Create)
		apiGroup.GET("scriptInfo/list", scriptInfo.List)
	}
}
