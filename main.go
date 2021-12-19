package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/config"
	"github.com/littlelittlelittleboy/scriptkilling/db"
	"github.com/littlelittlelittleboy/scriptkilling/route"
)

func main() {
	configure, err := config.Load("config.yaml")

	if err != nil {
		log.Printf("read config error, err {%s}", err)
		panic(err)
	}

	err = db.InitDB(configure.Mysql)
	if err != nil {
		log.Printf("db init error, err {%s}", err)
		panic(err)
	}

	// 1.创建路由
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("templates/*")
	group := r.Group("")
	route.RouterApp.Init(group)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
