package Request

import "main.go/extend/MirAi/v1/model/GroupRequestModel"

func BotInvitedJoinGroupRequest(qq string, json map[string]interface{}) {
	go GroupRequestModel.Api_insert(qq, json["eventId"], json["message"], json["fromId"], json["groupId"], json["groupName"], json["nick"])
}
