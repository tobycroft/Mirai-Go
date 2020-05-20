package route

import (
	"github.com/gin-gonic/gin"
	user2 "main.go/route/user"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	user := router.Group("/user")
	{
		user.Any("/", func(context *gin.Context) {
			context.String(0, user.BasePath())
		})
		v1 := user.Group("/v1")
		{
			user2.V1Router(v1)
		}
	}
}
