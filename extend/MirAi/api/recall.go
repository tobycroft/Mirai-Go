package api

import (
	"main.go/extend/MirAi/v1/action/Bot"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

/*
撤回消息
[POST] /recall
使用此方法撤回指定消息。对于bot发送的消息，有2分钟时间限制。对于撤回群聊中群员的消息，需要有相应权限

请求
{
    "sessionKey": "YourSession",
    "target": 987654321
}
名字	类型	可选	举例	说明
sessionKey	String	false	YourSession	已经激活的Session
target	Int	false	987654321	需要撤回的消息的messageId
响应: 返回统一状态码
{
    "code": 0,
    "msg": "success"
}
*/
func Recall(qq, messageId interface{}) (map[string]interface{}, error) {
	bot, err := Bot.BotSingle(qq)
	if err != nil {
		return nil, err
	} else {
		post := map[string]interface{}{
			"sessionKey": bot.SessionKey,
			"target":     messageId,
		}
		ret, err := Net.Postraw(bot.URL+"/recall", nil, post, nil, nil)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			return nil, err
		} else {
			return Jsong.JObject(ret)
		}
	}
}
