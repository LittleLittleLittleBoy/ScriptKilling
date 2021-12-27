package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/controller"
)

type LoginRouter struct{}

func (r *LoginRouter) Init(group *gin.RouterGroup) {
	webRouter := group.Group("")
	{
		webRouter.GET("login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})

	}

	apiRouter := group.Group("api")
	{
		user := controller.Instance.User
		apiRouter.POST("login", user.Login)
		apiRouter.POST("register", user.Register)
	}
}
