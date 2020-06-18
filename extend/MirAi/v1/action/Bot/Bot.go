package Bot

import (
	"main.go/extend/MirAi/v1/model/BotModel"
	"main.go/tuuz/Calc"
)

type Robot struct {
	QQ         string
	URL        string
	AuthKey    string
	SessionKey string
}

func BotSingle(qq interface{}) (Robot, bool) {
	bot := BotModel.Api_find(qq)
	if len(bot) < 1 {
		return Robot{}, false
	}
	url := Calc.Any2String(bot["url"])
	authKey := Calc.Any2String(bot["authKey"])
	sessionKey := Calc.Any2String(bot["sessionKey"])
	return Robot{URL: url, AuthKey: authKey, SessionKey: sessionKey, QQ: Calc.Any2String(qq)}, true
}

func BotAll() ([]Robot, bool) {
	bots := BotModel.Api_select()
	if len(bots) < 1 {
		return nil, false
	}
	Bots := []Robot{}
	for _, bot := range bots {
		qq := bot["qq"]
		url := Calc.Any2String(bot["url"])
		authKey := Calc.Any2String(bot["authKey"])
		sessionKey := Calc.Any2String(bot["sessionKey"])
		Bots = append(Bots, Robot{URL: url, AuthKey: authKey, SessionKey: sessionKey, QQ: Calc.Any2String(qq)})
	}
	return Bots, true
}
