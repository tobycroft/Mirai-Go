package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func MemberJoinRequestEvent(qq, eventId, fromId, groupId, operate, message interface{}) (map[string]interface{}, error) {
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
			"type":       "MemberJoinRequestEvent",
			"groupName":  "Group",
			"nick":       "Nick Name",
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
