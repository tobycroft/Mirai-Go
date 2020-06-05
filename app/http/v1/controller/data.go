package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DataController(route *gin.RouterGroup) {
	//route.Use(inital(), gin.Recovery())
	route.Any("index", index)
}

func index(c *gin.Context) {
	json, err := c.GetRawData()
	bot := c.GetHeader("bot")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bot)
		fmt.Println(string(json))
	}
}
