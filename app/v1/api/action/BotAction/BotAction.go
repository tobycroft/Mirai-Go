package BotAction

import "main.go/app/v1/api/model/BotFunctionModel"

func App_LockGroup(qq interface{}) bool {
	botfunction := BotFunctionModel.Api_find(qq)
	if len(botfunction) > 0 {
		if botfunction["lock_group"].(int64) == 1 {
			return true
		} else {
			return false
		}
	} else {
		BotFunctionModel.Api_insert(qq)
		return true
	}
}
