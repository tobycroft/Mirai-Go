package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/app/http/v1/action/MirRoute"
)

func DataController(route *gin.RouterGroup) {
	route.Use(inital, gin.Recovery())
	route.Any("", index)
}

func inital(c *gin.Context) {
	c.Next()
}

func index(c *gin.Context) {
	json, err := c.GetRawData()
	bot := c.GetHeader("bot")
	if err != nil {
		fmt.Println(err)
	} else {
		MirRoute.Do(bot, string(json))
	}
}
