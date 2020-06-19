package api

import (
	"errors"
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/extend/MirAi/v1/model/BotModel"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func Auth(qq interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
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

func App_Auth(qq interface{}) error {
	ret, err := Auth(qq)
	if err != nil {
		return err
	} else {
		if ret["code"].(float64) == 0 {
			session := ret["session"].(string)
			//bind
			ret, err = Verify(qq, session)
			if err != nil {
				return err
			} else {
				if BotModel.Api_update(qq, session) {
					return nil
				} else {
					return errors.New("数据库插入失败")
				}
			}
		} else {
			return errors.New("获取session失败")
		}
	}
}
