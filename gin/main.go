package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/gotools/gin/middleware"
	"github.com/lcsin/gotools/gin/routers"
)

func main() {
	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())

	// 路由组
	routers.Init(r)

	r.Run(":8080")
}
