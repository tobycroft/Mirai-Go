package http

import (
	"github.com/gin-gonic/gin"
	"main.go/app/http/v1/controller"
)

func V1Router(route *gin.RouterGroup) {
	controller.DataController(route.Group("/data"))
}
