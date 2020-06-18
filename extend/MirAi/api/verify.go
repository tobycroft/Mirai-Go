package api

import (
	"errors"
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func Verify(qq, sessionKey interface{}) (map[string]interface{}, error) {
	bot, ok := Bot.BotSingle(qq)
	if !ok {
		return nil, errors.New("未找到账号，可能机器人已经过期")
	}
	post := map[string]interface{}{
		"qq":         qq,
		"sessionKey": sessionKey,
	}
	ret, err := Net.Postraw(bot.URL+"/verify", nil, post, nil, nil)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
	} else {
		return Jsong.JObject(ret)
	}
	return nil, err
}
