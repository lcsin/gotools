package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/gotools/gin/message"
)

func AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, message.OK("add success"))
	return
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, message.OK(id))
}

func DelUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, message.OK(id))
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, message.OK(id))
}
