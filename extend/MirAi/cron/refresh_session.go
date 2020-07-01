package cron

import (
	"log"
	"main.go/extend/MirAi/api"
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

func Refresh_session() {
	for {
		Refresh()
		time.Sleep(10 * time.Minute)
	}
}

func Refresh() {
	bots, err := Bot.BotAll()
	if err != nil {
		log.Println("没有可用的bot")
	} else {
		for _, bot := range bots {
			qq := bot.QQ
			err := api.App_Auth(qq)
			if err != nil {
				Log.Crrs(err, tuuz.FUNCTION_ALL())
			}
		}
	}
}
