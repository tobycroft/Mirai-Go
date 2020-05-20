package dm

import (
	"github.com/gin-gonic/gin"
	"main.go/app/dm/v1/controller"
)

func V1Router(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	controller.ActController(route.Group("/act"))
	controller.ListController(route.Group("/list"))
}
