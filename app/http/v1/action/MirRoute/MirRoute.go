package MirRoute

import (
	"fmt"
	"main.go/app/http/v1/model/LogModel"
	"main.go/app/http/v1/model/LogUnknowModel"
	"main.go/extend/MirAi/v1/model/RequestGroupModel"
	"main.go/extend/MirAi/v1/model/RequestPrivateModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
)

func Do(qq, json string) {
	fmt.Println(qq)
	fmt.Println(json)
	go LogModel.Api_insert(qq, json)
	ret, err := Jsong.JObject(json)
	if err != nil {
		Log.Errs(err, tuuz.FUNCTION_ALL())
	} else {
		Type := Calc.Any2String(ret["type"])
		switch Type {

		case "FriendMessage":
			message(qq, Type, ret, json)
			break

		case "GroupNameChangeEvent":
			notice(qq, Type, ret, json)
			break

		case "NewFriendRequestEvent":
			request(qq, Type, ret, json)
			break

		default:
			break
		}
	}
}

func message(qq, Type string, json map[string]interface{}, str string) {

}

func notice(qq, Type string, json map[string]interface{}, str string) {

}

func request(qq, Type string, json map[string]interface{}, str string) {
	switch Type {
	case "NewFriendRequestEvent":
		go RequestPrivateModel.Api_insert(qq, json["eventId"], json["message"], json["fromId"], json["groupId"], json["nick"])
		break

	case "BotInvitedJoinGroupRequestEvent":
		go RequestGroupModel.Api_insert(qq, json["eventId"], json["message"], json["fromId"], json["groupId"], json["groupName"], json["nick"])
		break

	default:
		go LogUnknowModel.Api_insert(qq, str)
		break
	}
}
