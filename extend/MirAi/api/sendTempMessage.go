package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func SendTempMessage(qq, target interface{}, Type string, messages []interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey":   bot.SessionKey,
			"qq":           target,
			"group":        target,
			"messageChain": messages,
		}
		ret, err := Net.Postraw(bot.URL+"/sendTempMessage", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			return nil, err
		} else {
			return Jsong.JObject(ret)
		}
	}
}
