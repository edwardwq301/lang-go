package main

import (
	"gin-ex/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", handler.HelloWorld)
	r.GET("/ping", handler.Ping)

	r.Run(":8899")
}
