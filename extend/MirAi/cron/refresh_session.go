package cron

import (
	"errors"
	"fmt"
	"log"
	"main.go/extend/MirAi/api"
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/extend/MirAi/v1/model/BotModel"
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
			ret, err := api.Auth(qq)
			if err != nil {
				Log.Crrs(err, tuuz.FUNCTION_ALL())
			} else {
				if ret["code"].(float64) == 0 {
					session := ret["session"].(string)
					//bind
					ret, err = api.Verify(qq, session)
					if err != nil {
						Log.Crrs(err, tuuz.FUNCTION_ALL())
					} else {
						if BotModel.Api_update(qq, session) {
							fmt.Println(qq, "bot更新成功", session)
						} else {
							Log.Crrs(errors.New("bot_session更新失败,500"), tuuz.FUNCTION_ALL())
						}
					}
				} else {
					Log.Crrs(errors.New("获取session失败"), tuuz.FUNCTION_ALL())
				}
			}
		}
	}
}
