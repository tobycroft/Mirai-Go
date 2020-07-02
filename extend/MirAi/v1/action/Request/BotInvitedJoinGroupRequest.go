package Request

import (
	"main.go/app/v1/api/model/BotAllowGroupModel"
	"main.go/extend/MirAi/api"
	"main.go/extend/MirAi/v1/model/GroupRequestModel"
)

func BotInvitedJoinGroupRequest(qq string, json map[string]interface{}) {
	GroupRequestModel.Api_insert(qq, json["eventId"], json["message"], json["fromId"], json["groupId"], json["groupName"], json["nick"])
	allow := BotAllowGroupModel.Api_find(qq, json["groupId"])
	if len(allow) > 0 {
		api.MemberJoinRequestEvent(qq, json["eventId"], json["fromId"], json["groupId"], 0, "")
	} else {
		api.MemberJoinRequestEvent(qq, json["eventId"], json["fromId"], json["groupId"], 1, "机器人主不允许加入本群")
	}

}
