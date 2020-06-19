package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func About(qq interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	}
	ret, err := Net.Get(bot.URL+"/about", nil, nil, nil)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
	} else {
		return Jsong.JObject(ret)
	}
	return nil, err
}
