package http

import (
	"github.com/gin-gonic/gin"
	"main.go/app/http/v1/controller"
)

func V1Router(route *gin.RouterGroup) {
	//route.Any("/", func(context *gin.Context) {
	//	context.String(0, route.BasePath())
	//})
	controller.DataController(route.Group("/data"))

}
