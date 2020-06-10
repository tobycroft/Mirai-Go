package Message

import (
	"fmt"
	"main.go/app/http/v1/model/LogErrModel"
	"main.go/extend/MirAi/v1/model/GroupMessageModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"strings"
)

func GroupMessage(qq string, user_id int64, str string, sender map[string]interface{}, message_jsons []interface{}, message_id int64, messages string, imgs []string, time int64) {
	nickname := Calc.Any2String(sender["memberName"])
	user_role := Calc.Any2String(sender["permission"])
	group, err := Jsong.ParseObject(sender["group"])
	if err != nil {
		LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
		return
	}
	group_id, err := Calc.Any2Int64_2(group["id"])
	if err != nil {
		LogErrModel.Api_insert(err, tuuz.FUNCTION_ALL())
		return
	}
	role := Calc.Any2String(group["permission"])
	group_name := Calc.Any2String(group["name"])
	at := false
	var at_qq float64
	for _, txt := range message_jsons {
		msg, err := Jsong.ParseObject(txt)
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
		} else {
			switch msg["type"].(string) {
			case "Source":
				message_id = Calc.Any2Int64(msg["id"])
				time = Calc.Any2Int64(msg["time"])
				break

			case "Plain":
				messages += Calc.Any2String(msg["text"])
				break

			case "Image":
				imgs = append(imgs, Calc.Any2String(msg["url"]))
				break

			case "At":
				at_qq, err = Calc.Any2Float64_2(msg["target"])
				//todo：这里需要查看at的用户是否为机器人账号
				if err != nil {
					at = false
				} else {
					at = true
				}
				break

			case "AtAll":
				break

			case "Face":
				break

			case "Xml":
				break

			case "Json":
				break

			case "App":
				break

			default:
				break
			}
		}
	}
	if at {
		//todo:机器人被At
		fmt.Println(at, at_qq)
	}
	GroupMessageModel.Api_insert(qq, message_id, user_id, role, nickname, user_role, group_id, group_name, messages, strings.Join(imgs, ","), str, time)
}
