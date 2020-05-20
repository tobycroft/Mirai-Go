package up

import (
	"github.com/gin-gonic/gin"
	"main.go/app/up/v1/controller"
)

func V1Router(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})
	controller.RankController(route.Group("/rank"))
	controller.RoomController(route.Group("/room"))
	controller.UpcenterController(route.Group("/upcenter"))
	controller.LiveController(route.Group("/live"))
}
