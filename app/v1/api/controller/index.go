package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
}

func index(c *gin.Context) {
	c.JSON(200, RET.Ret_succ(0, index))
}
