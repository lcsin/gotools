package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(router gin.IRouter) {
	v1 := router.Group("/api/v1")

	// 静态资源
	v1.StaticFS("/file", http.Dir(""))
	// 健康检查
	v1.GET("health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 注册业务路由
	registerUserRouter(v1, "/user")
}

func registerUserRouter(router gin.IRouter, path string) {
	user := router.Group(path)

	user.POST("/add", AddUser)
	user.GET("/:id", GetUser)
	user.DELETE("/:id", DelUser)
	user.PUT("/:id", UpdateUser)
}
