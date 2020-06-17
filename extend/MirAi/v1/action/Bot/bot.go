package Bot

import (
	"main.go/extend/MirAi/v1/model/BotModel"
	"main.go/tuuz/Calc"
)

type Robot struct {
	URL        string
	AuthKey    string
	SessionKey string
}

func BotOne(qq string) Robot {
	bot := BotModel.Api_find(qq)
	url := Calc.Any2String(bot["url"])
	authKey := Calc.Any2String(bot["authKey"])
	sessionKey := Calc.Any2String(bot["sessionKey"])
	return Robot{URL: url, AuthKey: authKey, SessionKey: sessionKey}
}

func BotAll(qq string) []Robot {
	bots := BotModel.Api_select()
	Bots := []Robot{}
	for _, bot := range bots {
		url := Calc.Any2String(bot["url"])
		authKey := Calc.Any2String(bot["authKey"])
		sessionKey := Calc.Any2String(bot["sessionKey"])
		Bots = append(Bots, Robot{URL: url, AuthKey: authKey, SessionKey: sessionKey})
	}
	return Bots
}
