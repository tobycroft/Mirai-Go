package Bot

import (
	"main.go/extend/MirAi/v1/model/BotModel"
	"main.go/tuuz/Calc"
)

func All() {

}

type Bot struct {
	URL        string
	AuthKey    string
	SessionKey string
}

func Bota(qq string) Bot {
	bot := BotModel.Api_find(qq)
	url := Calc.Any2String(bot["url"])
	authKey := Calc.Any2String(bot["authKey"])
	sessionKey := Calc.Any2String(bot["sessionKey"])
	return Bot{URL: url, AuthKey: authKey, SessionKey: sessionKey}
}

func Bots(qq string) []Bot {
	bots := BotModel.Api_select()
	Bots := []Bot{}
	for _, bot := range bots {
		url := Calc.Any2String(bot["url"])
		authKey := Calc.Any2String(bot["authKey"])
		sessionKey := Calc.Any2String(bot["sessionKey"])
		Bots = append(Bots, Bot{URL: url, AuthKey: authKey, SessionKey: sessionKey})
	}

	return Bots
}
