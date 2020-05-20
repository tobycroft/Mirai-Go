package admin

import (
	"github.com/gin-gonic/gin"
	"main.go/app/admin/v1/controller"
)

func V1Router(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	controller.IndexController(route.Group("index"))
	controller.UserController(route.Group("user"))
	controller.RoomController(route.Group("room"))
	controller.AdminController(route.Group("admin"))
}
