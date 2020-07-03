package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

/*
退出群聊
使用此方法使Bot退出群聊

[POST] /quit
请求:
{
    "sessionKey": "YourSessionKey",
    "target": 123456789,
}
名字	可选	类型	举例	说明
sessionKey	false	String	"YourSessionKey"	你的session key
target	false	Long	123456789	退出的群号
响应
响应: 返回统一状态码
{
    "code": 0,
    "msg": "success"
}
bot为该群群主时退出失败并返回code 10(无操作权限)
*/
func Quit(qq, group_id interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     group_id,
		}
		ret, err := Net.Postraw(bot.URL+"/quit", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		} else {
			return Jsong.JObject(ret)
		}
		return nil, err
	}
}
