package route

import (
	"github.com/gin-gonic/gin"
	user2 "main.go/route/http"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	http := router.Group("/http")
	{
		http.Any("/", func(context *gin.Context) {
			context.String(0, router.BasePath())
		})
		v1 := http.Group("/v1")
		{
			user2.V1Router(v1)
		}
	}
}
