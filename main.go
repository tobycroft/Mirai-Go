package main

import (
	"github.com/gin-gonic/gin"
	"main.go/extend/MirAi/cron"
	"main.go/route"
)

func main() {
	go cron.Refresh_session()

	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	route.OnRoute(mainroute)
	mainroute.Run(":90")
}
