package Request

import "main.go/extend/MirAi/v1/model/PrivateRequestModel"

func NewFriendRequest(qq string, json map[string]interface{}) {
	go PrivateRequestModel.Api_insert(qq, json["eventId"], json["message"], json["fromId"], json["groupId"], json["nick"])
}
