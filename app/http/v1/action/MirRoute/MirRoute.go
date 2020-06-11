package MirRoute

import (
	"fmt"
	"main.go/app/http/v1/model/LogErrModel"
	"main.go/app/http/v1/model/LogModel"
	"main.go/app/http/v1/model/LogRecvModel"
	"main.go/app/http/v1/model/LogUnknowModel"
	"main.go/extend/MirAi/v1/action/Message"
	"main.go/extend/MirAi/v1/action/Request"
	"main.go/extend/MirAi/v1/model/GroupInfoModel"
	"main.go/extend/MirAi/v1/model/GroupMemberModel"
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
			message(qq, Type, ret, json)
			break

		case "GroupNameChangeEvent", "BotLeaveEventActive", "MemberJoinRequestEvent", "MemberJoinEvent",
			"GroupRecallEvent", "MemberLeaveEventQuit", "MemberLeaveEventKick", "BotGroupPermissionChangeEvent",
			"MemberMuteEvent", "MemberUnmuteEvent", "BotMuteEvent", "BotUnmuteEvent", "MemberCardChangeEvent",
			"MemberPermissionChangeEvent", "GroupMuteAllEvent":
			notice(qq, Type, ret, json)
			break

		case "NewFriendRequestEvent", "BotInvitedJoinGroupRequestEvent":
			request(qq, Type, ret, json)
			break

		default:
			break
		}
	}
}

func message(qq, Type string, json map[string]interface{}, str string) {
	sender, err := Jsong.ParseObject(json["sender"])
	if err != nil {
		LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
	} else {
		user_id, err := Calc.Any2Int64_2(sender["id"])
		if err != nil {
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		message_jsons, err := Jsong.ParseSlice(json["messageChain"])
		if err != nil {
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
		} else {
			go LogRecvModel.Api_insert(qq, str)
			var message_id int64 = 0 //消息的id为了避免多个机器人冲突，主库需要对qq和mid进行对应
			var messages = ""        //用于存储临时的消息去掉图片后的分析数据
			var time int64           //接收消息的准确时间
			var imgs []string        //img多个图片存成slice，没什么用以后做AI鉴黄要用

			switch Type {
			case "FriendMessage": //个人消息
				go Message.FriendMessage(qq, user_id, str, sender, message_jsons, message_id, messages, imgs, time)
				break

			case "GroupMessage": //群消息
				go Message.GroupMessage(qq, user_id, str, sender, message_jsons, message_id, messages, imgs, time)
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

func notice(qq, Type string, json map[string]interface{}, str string) {

	switch Type {
	case "GroupNameChangeEvent": //群-event-修改群名称
		member, err := Jsong.ParseObject(json["member"])
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		group, err := Jsong.ParseObject(member["group"])
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		GroupInfoModel.Api_update(qq, group["id"], member["new"], group["permission"])
		break

	case "BotLeaveEventActive": //群-event-机器人被T/群-被解散
		member, err := Jsong.ParseObject(json["member"])
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		group, err := Jsong.ParseObject(member["group"])
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		GroupInfoModel.Api_delete(qq, group["id"])
		break

	case "MemberJoinRequestEvent": //群-机器人-加入群
		if Calc.Any2String(json["fromId"]) == qq {
			GroupInfoModel.Api_insert(qq, json["groupId"], json["groupName"], "MEMBER")
		}
		break

	case "MemberJoinEvent": //群-用户-有人加群

		break

	case "GroupRecallEvent": //群-event-撤回
		break

	case "MemberLeaveEventQuit": //群-event-主动退群
		member, err := Jsong.ParseObject(json["member"])
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		group, err := Jsong.ParseObject(member["group"])
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		GroupMemberModel.Api_delete(group["id"], member["id"])
		break

	case "MemberLeaveEventKick": //群-用户-被T出群
		member, err := Jsong.ParseObject(json["member"])
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		group, err := Jsong.ParseObject(member["group"])
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
			return
		}
		GroupMemberModel.Api_delete(group["id"], member["id"])
		operator, err := Jsong.ParseObject(json["operator"])
		if err != nil {

		} else {
			//检测t人操作
			fmt.Println(operator)
		}
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
		LogUnknowModel.Api_insert(qq, str)
		break
	}
}

func request(qq, Type string, json map[string]interface{}, str string) {
	switch Type {
	case "NewFriendRequestEvent": //个人-event-收到好友申请
		go Request.NewFriendRequest(qq, json)
		break

	case "BotInvitedJoinGroupRequestEvent": //群-event-机器人被邀请进群
		go Request.BotInvitedJoinGroupRequest(qq, json)
		break

	default:
		go LogUnknowModel.Api_insert(qq, str)
		break
	}
}
