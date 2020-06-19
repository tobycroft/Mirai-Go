package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func sendImageMessage(qq interface{}, post map[string]interface{}, urls []string) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		ret, err := Net.Postraw(bot.URL+"/sendFriendMessage", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			return nil, err
		} else {
			return Jsong.JObject(ret)
		}
	}
}

func SendImageFriend(qq, user_id interface{}, urls []string) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     user_id,
			"urls":       urls,
		}
		return sendImageMessage(qq, post, urls)
	}
}

func SendImageGroup(qq, group_id interface{}, urls []string) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     group_id,
			"urls":       urls,
		}
		return sendImageMessage(qq, post, urls)
	}
}

func SendImageTemp(qq, group_id, user_id interface{}, urls []string) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"qq":         user_id,
			"group":      group_id,
			"urls":       urls,
		}
		return sendImageMessage(qq, post, urls)
	}
}
