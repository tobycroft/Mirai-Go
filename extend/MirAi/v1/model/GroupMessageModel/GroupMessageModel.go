package GroupMessageModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_message"

func Api_insert(qq, message_id, user_id, role, nickname, user_role, group_id, group_name, message, imgs, json, time interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"qq":         qq,
		"message_id": message_id,
		"user_id":    user_id,
		"role":       role,
		"nickname":   nickname,
		"user_role":  user_role,
		"group_id":   group_id,
		"group_name": group_name,
		"message":    message,
		"imgs":       imgs,
		"json":       json,
		"time":       time,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
