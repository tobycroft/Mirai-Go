package Message

import (
	"main.go/extend/MirAi/v1/model/PrivateMessageModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"strings"
)

func FriendMessage(qq *string, user_id *int64, str *string, sender map[string]interface{}, message_jsons []interface{}, message_id *int64, messages *string, imgs *[]string, time *int64) {
	nickname := Calc.Any2String(sender["nickname"])
	remark := Calc.Any2String(sender["remark"])

	for _, txt := range message_jsons {
		msg, err := Jsong.ParseObject(txt)
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
		} else {
			switch msg["type"].(string) {
			case "Source":
				*message_id = Calc.Any2Int64(msg["id"])
				*time = Calc.Any2Int64(msg["time"])
				break

			case "Plain":
				*messages += Calc.Any2String(msg["text"])
				break

			case "Image":
				*imgs = append(*imgs, Calc.Any2String(msg["Image"]))
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

			PrivateMessageModel.Api_insert(*qq, *message_id, *user_id, nickname, remark, *messages, strings.Join(*imgs, ","), *str, *time)
		}
	}
}
