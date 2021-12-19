package route

import (
	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/common/response"
)

type AdminRouter struct{}

func (r *AdminRouter) Init(group *gin.RouterGroup) {
	router := group.Group("/api/admin")
	router.GET("create", func(c *gin.Context) {
		response.GenerResponse(c, "ok")
	})
}
