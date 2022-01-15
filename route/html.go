package route

import (
	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/constants"
)

type HTMLRouter struct{}

func (r *HTMLRouter) Init(group *gin.RouterGroup) {
	group.Static("web", constants.HTML_BASE_PATH)
}
