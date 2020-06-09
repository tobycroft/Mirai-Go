package MirRoute

import (
	"fmt"
	"main.go/app/http/v1/model/LogErrModel"
	"main.go/app/http/v1/model/LogModel"
	"main.go/app/http/v1/model/LogRecvModel"
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

		case "FriendMessage", "GroupMessage", "TempMessage":
			message(&qq, &Type, ret, &json)
			break

		case "GroupNameChangeEvent", "BotLeaveEventActive", "MemberJoinRequestEvent", "MemberJoinEvent",
			"GroupRecallEvent", "MemberLeaveEventQuit", "MemberLeaveEventKick", "BotGroupPermissionChangeEvent",
			"MemberMuteEvent", "MemberUnmuteEvent", "BotMuteEvent", "BotUnmuteEvent", "MemberCardChangeEvent",
			"MemberPermissionChangeEvent", "GroupMuteAllEvent":
			notice(&qq, &Type, ret, &json)
			break

		case "NewFriendRequestEvent", "BotInvitedJoinGroupRequestEvent":
			request(&qq, &Type, ret, &json)
			break

		default:
			break
		}
	}
}

func message(qq, Type *string, json map[string]interface{}, str *string) {
	sender, err := Jsong.ParseObject(json["sender"])
	if err != nil {
		LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
	} else {
		messageChain, err := Jsong.ParseSlice(json["messageChain"])
		if err != nil {
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
		} else {
			go LogRecvModel.Api_insert(qq, str)
			switch *Type {
			case "FriendMessage": //个人消息
				if messageChain[0] != nil {

				}
				break

			case "GroupMessage": //群消息
				break

			case "TempMessage":
				break

			default:
				go LogUnknowModel.Api_insert(qq, str)
				break
			}
		}
	}
}

func notice(qq, Type *string, json map[string]interface{}, str *string) {
	switch *Type {
	case "GroupNameChangeEvent": //群-event-修改群名称
		break

	case "BotLeaveEventActive": //群-event-机器人被T/群-被解散
		break

	case "MemberJoinRequestEvent": //群-机器人-加入群
		break

	case "MemberJoinEvent": //群-用户-有人加群
		break

	case "GroupRecallEvent": //群-event-撤回
		break

	case "MemberLeaveEventQuit": //群-event-主动退群
		break

	case "MemberLeaveEventKick": //群-用户-被T出群
		break

	case "BotGroupPermissionChangeEvent": //群-机器人-被设定为管理员/群-机器人-被取消管理
		break

	case "MemberMuteEvent": //群-禁言-设定禁言1小时
		break

	case "MemberUnmuteEvent": //群-禁言-解除禁言
		break

	case "BotMuteEvent": //群-机器人-禁言
		break

	case "BotUnmuteEvent": //群-机器人-解除禁言
		break

	case "MemberCardChangeEvent": //群-机器人-被修改群名片
		break

	case "MemberPermissionChangeEvent": //群-管理-设为管理员/群-管理-取消管理员
		break

	case "GroupMuteAllEvent": //群-全员禁言-开/群-全员禁言-关
		break

	default:
		go LogUnknowModel.Api_insert(qq, str)
		break
	}
}

func request(qq, Type *string, json map[string]interface{}, str *string) {
	switch *Type {
	case "NewFriendRequestEvent": //个人-event-收到好友申请
		go RequestPrivateModel.Api_insert(qq, json["eventId"], json["message"], json["fromId"], json["groupId"], json["nick"])
		break

	case "BotInvitedJoinGroupRequestEvent": //群-event-机器人被邀请进群
		go RequestGroupModel.Api_insert(qq, json["eventId"], json["message"], json["fromId"], json["groupId"], json["groupName"], json["nick"])
		break

	default:
		go LogUnknowModel.Api_insert(qq, str)
		break
	}
}
