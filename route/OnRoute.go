package route

import (
	"github.com/gin-gonic/gin"
	http2 "main.go/route/http"
	v12 "main.go/route/v1"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	http := router.Group("http")
	{
		http.Any("/", func(context *gin.Context) {
			context.String(0, router.BasePath())
		})
		v1 := http.Group("v1")
		{
			http2.V1Router(v1)
		}
	}
	v1 := router.Group("v1")
	{
		v1.Any("/", func(context *gin.Context) {
			context.String(0, router.BasePath())
		})
		api := v1.Group("api")
		{
			v12.ApiRouter(api)
		}
	}
}
