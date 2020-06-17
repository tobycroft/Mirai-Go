package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/api/controller"
)

func ApiRouter(route *gin.RouterGroup) {
	controller.IndexController(route.Group("index"))
}
