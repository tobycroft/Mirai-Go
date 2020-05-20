package user

import (
	"github.com/gin-gonic/gin"
	"main.go/app/user/v1/controller"
)

func V1Router(route *gin.RouterGroup) {
	//route.Any("/", func(context *gin.Context) {
	//	context.String(0, route.BasePath())
	//})
	controller.UserController(route.Group("/user"))
	controller.FollowController(route.Group("/follow"))

}
