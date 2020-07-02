package BotAction

import "main.go/app/v1/api/model/BotFunctionModel"

func App_LockGroup(qq interface{}) int64 {
	botfunction := BotFunctionModel.Api_find(qq)
	if len(botfunction) > 0 {
		return botfunction["lock_group"].(int64)
	} else {
		BotFunctionModel.Api_insert(qq)
		return 1
	}
}
