package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func MuteAll(qq, target interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     target,
		}
		ret, err := Net.Postraw(bot.URL+"/muteAll", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		} else {
			return Jsong.JObject(ret)
		}
		return nil, err
	}
}

func UnmuteAll(qq, target interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     target,
		}
		ret, err := Net.Postraw(bot.URL+"/unmuteAll", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		} else {
			return Jsong.JObject(ret)
		}
		return nil, err
	}
}

func Mute(qq, group_id, user_id, duration interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     group_id,
			"memberId":   user_id,
			"time":       duration,
		}
		ret, err := Net.Postraw(bot.URL+"/mute", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		} else {
			return Jsong.JObject(ret)
		}
		return nil, err
	}
}
func Unmute(qq, group_id, user_id, duration interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     group_id,
			"memberId":   user_id,
		}
		ret, err := Net.Postraw(bot.URL+"/unmute", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		} else {
			return Jsong.JObject(ret)
		}
		return nil, err
	}
}
