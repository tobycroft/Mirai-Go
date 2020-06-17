package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/extend/MirAi/api"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
	route.Any("about", about)
	route.Any("auth", auth)
}

func index(c *gin.Context) {
	c.JSON(200, RET.Ret_succ(0, index))
}

func about(c *gin.Context) {
	ret, err := api.About(2140300010)
	if err != nil {
		c.JSON(200, err.Error())
	} else {
		c.JSON(200, ret)
	}
}

func auth(c *gin.Context) {
	ret, err := api.Auth(2140300010)
	if err != nil {
		c.JSON(200, err.Error())
	} else {
		c.JSON(200, ret)
	}
}

func verify(c *gin.Context) {
}
