package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func MemberList(qq, target interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		get := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     target,
		}
		ret, err := Net.Get(bot.URL+"/groupList", get, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		} else {
			return Jsong.JObject(ret)
		}
		return nil, err
	}
}
