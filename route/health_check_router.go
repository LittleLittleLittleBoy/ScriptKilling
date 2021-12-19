package route

import (
	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/common/response"
)

type HealthCheckRouter struct{}

func (r *HealthCheckRouter) Init(group *gin.RouterGroup) {
	router := group.Group("api")
	router.GET("health", func(c *gin.Context) {
		response.GenerResponse(c, "ok")
	})
}
