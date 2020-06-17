package api

import (
	"errors"
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func About(qq interface{}) (map[string]interface{}, error) {
	bot, ok := Bot.BotSingle(qq)
	if !ok {
		return nil, errors.New("未找到账号，可能机器人已经过期")
	}
	ret, err := Net.Get(bot.URL+"/about", nil, nil, nil)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
	} else {
		return Jsong.JObject(ret)
	}
	return nil, err
}

func Auth(qq interface{}) (map[string]interface{}, error) {
	bot, ok := Bot.BotSingle(qq)
	if !ok {
		return nil, errors.New("未找到账号，可能机器人已经过期")
	}
	post := map[string]interface{}{
		"authKey": bot.AuthKey,
	}
	ret, err := Net.Postraw(bot.URL+"/auth", nil, post, nil, nil)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
	} else {
		return Jsong.JObject(ret)
	}
	return nil, err
}

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

func Release(qq, sessionKey interface{}) (map[string]interface{}, error) {
	bot, ok := Bot.BotSingle(qq)
	if !ok {
		return nil, errors.New("未找到账号，可能机器人已经过期")
	}
	post := map[string]interface{}{
		"qq":         qq,
		"sessionKey": sessionKey,
	}
	ret, err := Net.Postraw(bot.URL+"/release", nil, post, nil, nil)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
	} else {
		return Jsong.JObject(ret)
	}
	return nil, err
}
