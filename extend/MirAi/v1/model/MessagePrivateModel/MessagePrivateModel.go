package MessagePrivateModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "message_private"

func Api_insert(qq, message_id, user_id, nickname, remark, message, imgs, json, time interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"qq":         qq,
		"message_id": message_id,
		"user_id":    user_id,
		"nickname":   nickname,
		"remark":     remark,
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
