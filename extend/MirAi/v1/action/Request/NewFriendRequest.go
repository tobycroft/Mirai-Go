package Request

import (
	"main.go/app/v1/api/action/BotAction"
	"main.go/app/v1/api/model/BotAllowGroupModel"
	"main.go/extend/MirAi/api"
	"main.go/extend/MirAi/v1/model/PrivateRequestModel"
)

func NewFriendRequest(qq string, json map[string]interface{}) {
	PrivateRequestModel.Api_insert(qq, json["eventId"], json["message"], json["fromId"], json["groupId"], json["nick"])
	if BotAction.App_lockFriend(qq) {
		allow := BotAllowGroupModel.Api_find(qq, json["groupId"])
		if len(allow) > 0 {
			api.NewFriendRequestEvent(qq, json["eventId"], json["fromId"], json["groupId"], 0, "")
		} else {
			api.MemberJoinRequestEvent(qq, json["eventId"], json["fromId"], json["groupId"], 1, "不能添加，请增加列表")
		}
	} else {
		api.MemberJoinRequestEvent(qq, json["eventId"], json["fromId"], json["groupId"], 0, "")
	}
}
