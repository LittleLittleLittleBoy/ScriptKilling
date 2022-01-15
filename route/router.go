package route

import "github.com/gin-gonic/gin"

type InitRouter interface {
	Init(group *gin.RouterGroup)
}

type RouterGroup struct {
	routerList []InitRouter
}

func (r RouterGroup) Init(group *gin.RouterGroup) {
	for _, router := range r.routerList {
		router.Init(group)
	}
}

var RouterApp = RouterGroup{
	routerList: []InitRouter{
		&HealthCheckRouter{},
		&ScriptCreatorRouter{},
		&LoginRouter{},
		&HTMLRouter{},
	},
}
