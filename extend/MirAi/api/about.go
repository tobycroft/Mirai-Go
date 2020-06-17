package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func About(qq string) (map[string]interface{}, error) {
	bot := Bot.BotSingle(qq)
	ret, err := Net.Get(bot.URL+"/about", nil, nil, nil)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
	} else {
		return Jsong.ParseObject(ret)
	}
	return nil, err
}

func Auth(qq string) (map[string]interface{}, error) {
	bot := Bot.BotSingle(qq)
	post := map[string]interface{}{
		"authKey": bot.AuthKey,
	}
	Net.Postraw(bot.URL+"/auth", nil, post, nil, nil)
}
