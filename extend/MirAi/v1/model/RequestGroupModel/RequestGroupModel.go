package RequestGroupModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "request_group"

func Api_insert(qq, eventId, message, fromId, groupId, groupName, nick interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"qq":        qq,
		"eventId":   eventId,
		"message":   message,
		"fromId":    fromId,
		"groupId":   groupId,
		"groupName": groupName,
		"nick":      nick,
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
