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
		refresh()
		time.Sleep(10 * time.Minute)
	}
}

func refresh() {
	bots, ok := Bot.BotAll()
	if !ok {
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
