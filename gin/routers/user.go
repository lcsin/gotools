package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/gotools/gin/message"
)

type UserParams struct {
	ID   int64  `json:"id" binding:"required,gt=0"`
	Name string `json:"name" binding:"required"`
	Age  string `json:"age" binding:"required"`
}

// AddUser 添加用户
// @POST /api/v1/user
// @Permission [p1,p2,p3...]
func AddUser(c *gin.Context) {
	var param *UserParams
	if err := c.ShouldBind(param); err != nil {
		c.JSON(http.StatusOK, message.Error(http.StatusBadRequest, err.Error()))
		return
	}

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
