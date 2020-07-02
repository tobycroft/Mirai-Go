package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func AllowJoinIn(qq, eventId, fromId, groupId, operate, message interface{}) (map[string]interface{}, error) {
	/*
		operate	说明
		0	同意邀请
		1	拒绝邀请
	*/
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"eventId":    eventId,
			"fromId":     fromId,
			"groupId":    groupId,
			"operate":    operate,
			"message":    message,
		}
		ret, err := Net.Postraw(bot.URL+"/resp/memberJoinRequestEvent", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			return nil, err
		} else {
			return Jsong.JObject(ret)
		}
	}
}

func AllowInviteGroup(qq, eventId, fromId, groupId, operate, message interface{}) (map[string]interface{}, error) {
	/*
		operate	说明
		0	同意邀请
		1	拒绝邀请
	*/
	return AllowJoinIn(qq, eventId, fromId, groupId, operate, message)
}
