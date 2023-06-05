package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/gotools/gin/middleware"
	"github.com/lcsin/gotools/gin/routers"
)

func main() {
	engine := gin.Default()
	// 中间件
	engine.Use(middleware.Cors())
	// 路由组
	routers.Init(engine)

	engine.Run(":8080")
}
