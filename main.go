package main

import (
	"formForHumans/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	route:=gin.Default()

	route.GET("/", controller.GetHandler)
	route.POST("/post", controller.PostHandler)
	route.Run(controller.WEBSERVERPORT)
}
