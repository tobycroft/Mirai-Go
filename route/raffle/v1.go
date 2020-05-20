package raffle

import (
	"github.com/gin-gonic/gin"
	"main.go/app/raffle/v1/controller"
)

func V1Router(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	controller.ListController(route.Group("/list"))
	controller.BagController(route.Group("/bag"))
	controller.SendController(route.Group("send"))
}
