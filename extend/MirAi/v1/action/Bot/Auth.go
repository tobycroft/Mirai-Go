package Bot

import (
	"errors"
	"main.go/extend/MirAi/api"
	"main.go/extend/MirAi/v1/model/BotModel"
)

func App_auth(qq interface{}) error {
	ret, err := api.Auth(qq)
	if err != nil {
		return err
	} else {
		if ret["code"].(float64) == 0 {
			session := ret["session"].(string)
			//bind
			ret, err = api.Verify(qq, session)
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
